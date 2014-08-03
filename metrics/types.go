package metrics

import (
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
	data map[string]map[string]metric
}

func (store Storage) Get(family string, name string) int64 {
	metric := store.data[family][name]
	return metric.get()
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
}

type gauge struct {
	buf *util.RingBuffer
}

type counter struct {
	buf       *util.RingBuffer
	transform func(cur, prev int64) int64
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

func (store Storage) declareGauge(family string, name string) {
	fam := store.data[family]
	if fam == nil {
		store.data[family] = map[string]metric{}
		fam = store.data[family]
	}

	data := fam[name]
	if data == nil {
		fam[name] = gauge{util.NewRingBuffer(SLOTS)}
		data = fam[name]
	}
}

func (store Storage) declareCounter(family string, name string, xform func(cur, prev int64) int64) {
	fam := store.data[family]
	if fam == nil {
		store.data[family] = map[string]metric{}
		fam = store.data[family]
	}

	data := fam[name]
	if data == nil {
		fam[name] = counter{util.NewRingBuffer(SLOTS), xform}
		data = fam[name]
	}
}

func (store Storage) save(family string, name string, value int64) {
	m := store.data[family][name]
	if m == nil {
		panic(family + "/" + name)
	}
	m.put(value)
}

func (store Storage) saveType(family string, name string, value int64, t metricType) {
	m := store.data[family][name]
	if m == nil {
		// dynamically declare metrics for disk metrics where the name
		// is dynamic based on the mount point
		store.declareGauge(family, name)
		m = store.data[family][name]
	}
	m.put(value)
}
