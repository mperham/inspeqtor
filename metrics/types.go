package metrics

import (
	"errors"
	"github.com/mperham/inspeqtor/util"
	"regexp"
	"sort"
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

	// all metric ring buffers will store one hour of metric history
	SLOTS = 3600 / 15
)

type Type uint8

const (
	Counter Type = iota
	Gauge
)

var (
	meminfoParser = regexp.MustCompile("([^:]+):\\s+(\\d+)")
	swapRegexp    = regexp.MustCompile("= (\\d+\\.\\d{2}[A-Z])(.*)")
	multiplyBy100 = func(val float64) float64 {
		return val * 100
	}
	displayLoad = func(val float64) string {
		return strconv.FormatFloat(val, 'f', 2, 64)
	}
	DisplayPercent = func(val float64) string {
		return strconv.FormatFloat(val, 'f', 1, 64) + "%"
	}
	DisplayInMB = func(val float64) string {
		return strconv.FormatFloat(val/(1024*1024), 'f', 2, 64) + "m"
	}
)

// transform the raw collected data into something we can compare.  Used by cpu(*)
// to transform raw ticks into a percentage.
type TransformFunc func(float64, float64) float64

// Convert the raw metric value into something displayable to the user.
type DisplayFunc func(float64) string

type Store interface {
	Get(family string, name string) float64
	Display(family string, name string) string
	Collect(pid int) error

	Families() []string
	Metrics(family string) []string

	Save(family, name string, value float64)
	DeclareCounter(family, name string, xform TransformFunc, display DisplayFunc)
	DeclareGauge(family, name string, display DisplayFunc)
	Buffer(family, name string) *util.RingBuffer
}

type Loadable interface {
	Load(values ...interface{})
}

type storage struct {
	tree map[string]*family
}

func (store *storage) Buffer(family, name string) *util.RingBuffer {
	f := store.tree[family]
	if f == nil {
		return nil
	}

	m := f.metrics[name]
	if m == nil {
		return nil
	}

	return m.buffer()
}

func (store *storage) Families() []string {
	families := []string{}
	for k, _ := range store.tree {
		families = append(families, k)
	}
	sort.Strings(families)
	return families
}

func (store *storage) Metrics(family string) []string {
	met := []string{}
	for k, _ := range store.tree[family].metrics {
		met = append(met, k)
	}
	sort.Strings(met)
	return met
}

func (store *storage) Get(family string, name string) float64 {
	metric, _ := store.find(family, name)
	if metric == nil {
		// This can happen when using an Inspeqtor Pro .inq file
		// with Inspeqtor, since metrics like mysql(Queries) won't exist.
		util.Warn("BUG: Metric %s(%s) does not exist", family, name)
		return 0
	}
	return metric.get()
}

func (store *storage) Display(family string, name string) string {
	metric, _ := store.find(family, name)
	return metric.display()
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

type metric interface {
	put(val float64)
	get() float64
	display() string
	buffer() *util.RingBuffer
}

type gauge struct {
	buf        *util.RingBuffer
	forDisplay DisplayFunc
}

type counter struct {
	buf        *util.RingBuffer
	transform  TransformFunc
	forDisplay DisplayFunc
}

func (g *gauge) buffer() *util.RingBuffer {
	return g.buf
}

func (c *counter) buffer() *util.RingBuffer {
	return c.buf
}

func (g *gauge) put(val float64) {
	g.buf.Add(val)
}

func (c *counter) put(val float64) {
	c.buf.Add(val)
}

func (g *gauge) get() float64 {
	val := g.buf.At(0)
	if val == nil {
		return -1
	}
	return *val
}

func (g *gauge) display() string {
	val := g.get()
	if g.forDisplay != nil {
		return g.forDisplay(val)
	} else {
		return strconv.FormatFloat(val, 'f', 1, 64)
	}
}

func (c *counter) display() string {
	val := c.get()
	if c.forDisplay != nil {
		return c.forDisplay(val)
	} else {
		return strconv.FormatFloat(val, 'f', 1, 64)
	}
}

/*
 * Counter values should be monotonically increasing.
 * The value of a counter is actually the difference between two values.
 */
func (c *counter) get() float64 {
	cur := c.buf.At(0)
	prev := c.buf.At(-1)
	if cur == nil || prev == nil {
		return 0
	}
	if c.transform != nil {
		return c.transform(*cur, *prev)
	} else {
		return *cur - *prev
	}
}

func (store *storage) fill(values ...interface{}) {
	fam := values[0].(string)
	name := values[1].(string)
	for _, val := range values[2:] {
		store.Save(fam, name, float64(val.(int)))
	}
}

func (store *storage) declareDynamicFamily(familyName string) {
	store.tree[familyName] = &family{familyName, true, map[string]metric{}}
}

func (store *storage) DeclareGauge(familyName string, name string, display DisplayFunc) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = &gauge{util.NewRingBuffer(SLOTS), display}
		data = fam.metrics[name]
	}
}

func (store *storage) DeclareCounter(familyName string, name string, xform TransformFunc, display DisplayFunc) {
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

func (store *storage) Save(family string, name string, value float64) {
	m := store.tree[family].metrics[name]
	if m == nil {
		panic("No such metric: " + displayName(family, name))
	}
	m.put(value)
}

func (store *storage) saveType(family string, name string, value float64, t Type) {
	fam := store.tree[family]
	met := fam.metrics[name]
	if met == nil && fam.allowDynamic {
		// declare metrics for disk metrics where the name
		// is dynamic based on the mount point
		if t == Gauge {
			store.DeclareGauge(family, name, DisplayPercent)
		} else {
			store.DeclareCounter(family, name, nil, nil)
		}
		met = store.tree[family].metrics[name]
	}
	met.put(value)
}
