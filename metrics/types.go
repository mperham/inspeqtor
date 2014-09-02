package metrics

import (
	"errors"
	"inspeqtor/util"
	"strconv"
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

type Store interface {
	Get(family string, name string) int64
	Display(family string, name string) string
	PrepareRule(family string, name string, threshold int64) (int64, error)
	Collect(pid int) error
}

type storage struct {
	tree map[string]*family
}

func (store *storage) Get(family string, name string) int64 {
	metric, _ := store.find(family, name)
	return metric.get()
}

func (store *storage) Display(family string, name string) string {
	metric, _ := store.find(family, name)
	return metric.display()
}

func (store *storage) PrepareRule(family string, name string, threshold int64) (int64, error) {
	metric, err := store.find(family, name)
	if err != nil {
		return 0, err
	}
	if metric == nil {
		return threshold, nil
	}
	return metric.prepare(threshold), nil
}

func (store *storage) find(family, name string) (metric, error) {
	fam := store.tree[family]
	if fam == nil {
		return nil, nil
	}

	metric := fam.metrics[name]
	if metric == nil && !fam.allowDynamic {
		return nil, errors.New("No such metric: " + displayName(family, name))
	}
	return metric, nil
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

type prepareFunc func(int64) int64
type transformFunc func(int64, int64) int64
type displayFunc func(int64) string

type metric interface {
	put(val int64)
	get() int64
	display() string

	// Prepare is called on a rule threshold to ensure it's in the
	// same format as the collected metric values.
	// For example, load average is expressed as 1.55 but internally
	// stored as 155 since all metrics are int64.  We "prepare" a rule
	// threshold of 10 by multiplying it by 100 so the actual threshold
	// internally is 1000.
	prepare(threshold int64) int64
}

type gauge struct {
	buf           *util.RingBuffer
	prepThreshold prepareFunc
	forDisplay    displayFunc
}

type counter struct {
	buf        *util.RingBuffer
	transform  transformFunc
	forDisplay displayFunc
}

func (g *gauge) prepare(val int64) int64 {
	if g.prepThreshold != nil {
		return g.prepThreshold(val)
	} else {
		return val
	}
}

func (c *counter) prepare(val int64) int64 {
	return val
}

func (g *gauge) put(val int64) {
	g.buf.Add(val)
}

func (c *counter) put(val int64) {
	c.buf.Add(val)
}

func (g *gauge) get() int64 {
	cur := g.buf.At(0)
	if cur == nil {
		return -1
	}
	return cur.(int64)
}

func (g *gauge) display() string {
	val := g.get()
	if g.forDisplay != nil {
		return g.forDisplay(val)
	} else {
		return strconv.FormatInt(val, 10)
	}
}

func (c *counter) display() string {
	val := c.get()
	if c.forDisplay != nil {
		return c.forDisplay(val)
	} else {
		return strconv.FormatInt(val, 10)
	}
}

/*
 * Counter values should be monotonically increasing.
 * The value of a counter is actually the difference between two values.
 */
func (c *counter) get() int64 {
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

func (store *storage) fill(values ...interface{}) {
	fam := values[0].(string)
	name := values[1].(string)
	for _, val := range values[2:] {
		store.save(fam, name, int64(val.(int)))
	}
}

func (store *storage) declareDynamicFamily(familyName string) {
	store.tree[familyName] = &family{familyName, true, map[string]metric{}}
}

func (store *storage) declareGauge(familyName string, name string, prep prepareFunc, display displayFunc) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = &gauge{util.NewRingBuffer(SLOTS), prep, display}
		data = fam.metrics[name]
	}
}

func (store *storage) declareCounter(familyName string, name string, xform transformFunc, display displayFunc) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = &counter{util.NewRingBuffer(SLOTS), xform, display}
		data = fam.metrics[name]
	}
}

func (store *storage) save(family string, name string, value int64) {
	m := store.tree[family].metrics[name]
	if m == nil {
		panic("No such metric: " + displayName(family, name))
	}
	m.put(value)
}

func (store *storage) saveType(family string, name string, value int64, t metricType) {
	fam := store.tree[family]
	met := fam.metrics[name]
	if met == nil && fam.allowDynamic {
		// declare metrics for disk metrics where the name
		// is dynamic based on the mount point
		if t == Gauge {
			store.declareGauge(family, name, nil, displayPercent)
		} else {
			store.declareCounter(family, name, nil, nil)
		}
		met = store.tree[family].metrics[name]
	}
	met.put(value)
}
