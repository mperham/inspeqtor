package metrics

import (
	"errors"
	"fmt"
	"inspeqtor/util"
)

const (
	/*
	  CPU time is stored in system clock ticks.  Most
	  modern linux systems use 100 clock ticks per second,
	  or 100 Hz.  Use "getconf CLK_TCK" to verify this.
	  Since our cycle time is 15 seconds, there's 1500
	  ticks to spend each cycle.  If a process uses 750
	  ticks on the CPU, that means it used 50% of the CPU
	  during that cycle.  Multithreaded processes running on
	  systems with multiple CPUs/cores can use more than 100% CPU.
	*/
	CLK_TCK = 100
)

var (
	SLOTS = 3600 / 15
)

type Storage struct {
	tree map[string]*family
}

func (store Storage) Get(family string, name string) int64 {
	fam := store.tree[family]
	if fam == nil {
		panic(fmt.Sprintf("No such metric: %s", displayName(family, name)))
	}
	metric := fam.metrics[name]
	if metric == nil {
		panic(fmt.Sprintf("No such metric: %s", displayName(family, name)))
	}
	return metric.get()
}

func (store Storage) PrepareRule(family string, name string, threshold int64) (int64, error) {
	fam := store.tree[family]
	if fam == nil {
		return 0, errors.New("No such metric family: " + family)
	}

	metric := fam.metrics[name]
	if metric == nil && !fam.allowDynamic {
		return 0, errors.New("No such metric: " + displayName(family, name))
	}

	return metric.prepare(threshold), nil
}

func displayName(family, name string) string {
	s := family
	if name != "" {
		return s + "(" + name + ")"
	}
	return s
}

type family struct {
	name         string
	allowDynamic bool
	metrics      map[string]metric
}

// private

type metricType uint8

const (
	Counter metricType = iota
	Gauge
)

type metric interface {
	put(val int64)
	get() int64
	prepare(threshold int64) int64
}

type gauge struct {
	buf           *util.RingBuffer
	prepThreshold func(thresh int64) int64
}

type counter struct {
	buf           *util.RingBuffer
	transform     func(cur, prev int64) int64
	prepThreshold func(thresh int64) int64
}

func (g gauge) prepare(val int64) int64 {
	if g.prepThreshold != nil {
		return g.prepThreshold(val)
	} else {
		return val
	}
}

func (c counter) prepare(val int64) int64 {
	if c.prepThreshold != nil {
		return c.prepThreshold(val)
	} else {
		return val
	}
}

func (g gauge) put(val int64) {
	g.buf.Add(val)
}

func (c counter) put(val int64) {
	c.buf.Add(val)
}

func (g gauge) get() int64 {
	cur := g.buf.At(0)
	if cur == nil {
		return -1
	}
	return cur.(int64)
}

/*
 * Counter values should be monotonically increasing.
 * The value of a counter is actually the difference between two values.
 */
func (c counter) get() int64 {
	cur := c.buf.At(0)
	prev := c.buf.At(-1)
	if cur == nil || prev == nil {
		return 0
	}
	if c.transform != nil {
		return c.transform(cur.(int64), prev.(int64))
	} else {
		return cur.(int64) - prev.(int64)
	}
}

func (store Storage) fill(values ...interface{}) {
	fam := values[0].(string)
	name := values[1].(string)
	for _, val := range values[2:] {
		store.save(fam, name, int64(val.(int)))
	}
}

func (store Storage) declareDynamicFamily(familyName string) {
	store.tree[familyName] = &family{familyName, true, map[string]metric{}}
}

func (store Storage) declareGauge(familyName string, name string) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = gauge{util.NewRingBuffer(SLOTS), nil}
		data = fam.metrics[name]
	}
}

func (store Storage) declareCounter(familyName string, name string, xform func(cur, prev int64) int64) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = counter{util.NewRingBuffer(SLOTS), xform, nil}
		data = fam.metrics[name]
	}
}

func (store Storage) save(family string, name string, value int64) {
	m := store.tree[family].metrics[name]
	if m == nil {
		panic("No such metric: " + displayName(family, name))
	}
	m.put(value)
}

func (store Storage) saveType(family string, name string, value int64, t metricType) {
	fam := store.tree[family]
	met := fam.metrics[name]
	if met == nil && fam.allowDynamic {
		// dynamically declare metrics for disk metrics where the name
		// is dynamic based on the mount point
		if t == Gauge {
			store.declareGauge(family, name)
		} else {
			store.declareCounter(family, name, nil)
		}
		met = store.tree[family].metrics[name]
	}
	met.put(value)
}
