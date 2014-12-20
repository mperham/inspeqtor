package metrics

import (
	"errors"
	"regexp"
	"sort"
	"strconv"

	"github.com/mperham/inspeqtor/util"
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
	ClkTck = 100

	// all metric ring buffers will store one hour of metric history
	Slots = 3600 / 15
)

type SourceBuilder func(map[string]string) (Source, error)
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
	Sources = map[string]SourceBuilder{}
)

// transform the raw collected data into something we can compare.  Used by cpu:*
// to transform raw ticks into a percentage.
type TransformFunc func(float64, float64) float64

// Convert the raw metric value into something displayable to the user.
type DisplayFunc func(float64) string

type Map map[string]float64

type Descriptor struct {
	Name       string
	MetricType Type
	Display    DisplayFunc
	Transform  TransformFunc
}

type MandatorySource interface {
	Mandatory() bool
}

type Source interface {
	Name() string
	// Called once before any metrics are captured
	Prepare() error
	// Called every cycle to collect metrics
	Capture() (Map, error)
	// opt into watching this metric.
	// collectors don't have to collect every possible metric,
	// for efficiency reasons.
	Watch(metricName string)
	ValidMetrics() []Descriptor
}

type Store interface {
	Readable
	Writable
	Collectable
}

type Readable interface {
	Get(family string, name string) float64
	Display(family string, name string) string
	Families() []string
	MetricNames(family string) []string
	Metric(family, name string) Metric
	Each(func(family, name string, metric Metric))
}

type Collectable interface {
	AddSource(name string, config map[string]string) (Source, error)
	Prepare() error

	Collect(pid int) error
	// declare that a rule wants to act on this metric.
	// useful if we only want to collect a metric if a
	// rule will act upon it.
	Watch(family, name string) error
}

type Writable interface {
	Save(family, name string, value float64)
	DeclareCounter(family, name string, xform TransformFunc, display DisplayFunc)
	DeclareGauge(family, name string, display DisplayFunc)
}

type Loadable interface {
	Load(values ...interface{})
}

type storage struct {
	tree map[string]*family
}

func (store *storage) Each(iter func(family, name string, metric Metric)) {
	for _, fam := range store.Families() {
		for _, met := range store.MetricNames(fam) {
			iter(fam, met, store.Metric(fam, met))
		}
	}
}

func (store *storage) Metric(family, name string) Metric {
	f := store.tree[family]
	if f == nil {
		return nil
	}

	x, ok := f.metrics[name]
	if !ok {
		return nil
	}
	return x
}

func (store *storage) Families() []string {
	families := []string{}
	for k := range store.tree {
		families = append(families, k)
	}
	sort.Strings(families)
	return families
}

func (store *storage) MetricNames(family string) []string {
	met := []string{}
	fam := store.tree[family]
	if fam == nil {
		return met
	}

	for k := range fam.metrics {
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
	return metric.Get()
}

func (store *storage) Display(family string, name string) string {
	metric, _ := store.find(family, name)
	return metric.Display()
}

func (store *storage) find(family, name string) (Metric, error) {
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
	metrics      map[string]Metric
}

// private

type Metric interface {
	// Used for current value
	Put(val float64)
	Get() float64
	Display() string

	// Used for History ('show')
	At(int) *float64
	Displayable(float64) string
	Size() int
	Type() Type
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

func (g *gauge) Size() int {
	return g.buf.Size()
}

func (c *counter) Size() int {
	sz := c.buf.Size() - 1
	if sz < 0 {
		return 0
	}
	return sz
}

func (g *gauge) Type() Type {
	return Gauge
}

func (c *counter) Type() Type {
	return Counter
}

func (g *gauge) Put(val float64) {
	g.buf.Add(val)
}

func (c *counter) Put(val float64) {
	c.buf.Add(val)
}

func (g *gauge) Get() float64 {
	v := g.At(0)
	if v == nil {
		return -1
	}
	return *v
}

func (g *gauge) At(idx int) *float64 {
	return g.buf.At(idx)
}

func (g *gauge) Display() string {
	val := g.Get()
	return g.Displayable(val)
}

func (g *gauge) Displayable(val float64) string {
	if g.forDisplay != nil {
		return g.forDisplay(val)
	}
	return strconv.FormatFloat(val, 'f', 1, 64)
}

func (c *counter) Display() string {
	val := c.Get()
	return c.Displayable(val)
}

func (c *counter) Displayable(val float64) string {
	if c.forDisplay != nil {
		return c.forDisplay(val)
	}
	return strconv.FormatFloat(val, 'f', 1, 64)
}

/*
 * Counter values should be monotonically increasing.
 * The value of a counter is actually the difference between two values.
 */
func (c *counter) Get() float64 {
	v := c.At(0)
	if v == nil {
		return 0
	}
	return *v
}

func (c *counter) At(idx int) *float64 {
	cur := c.buf.At(idx)
	prev := c.buf.At(idx - 1)
	if cur == nil || prev == nil {
		return nil
	}
	var x float64
	if c.transform != nil {
		x = c.transform(*cur, *prev)
	} else {
		x = *cur - *prev
	}
	return &x
}

func (store *storage) fill(values ...interface{}) {
	fam := values[0].(string)
	name := values[1].(string)
	for _, val := range values[2:] {
		store.Save(fam, name, float64(val.(int)))
	}
}

func (store *storage) declareDynamicFamily(familyName string) {
	store.tree[familyName] = &family{familyName, true, map[string]Metric{}}
}

func (store *storage) DeclareGauge(familyName string, name string, display DisplayFunc) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]Metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = &gauge{util.NewRingBuffer(Slots), display}
		data = fam.metrics[name]
	}
}

func (store *storage) DeclareCounter(familyName string, name string, xform TransformFunc, display DisplayFunc) {
	fam := store.tree[familyName]
	if fam == nil {
		store.tree[familyName] = &family{familyName, false, map[string]Metric{}}
		fam = store.tree[familyName]
	}

	data := fam.metrics[name]
	if data == nil {
		fam.metrics[name] = &counter{util.NewRingBuffer(Slots), xform, display}
		data = fam.metrics[name]
	}
}

func (store *storage) Save(family string, name string, value float64) {
	f := store.tree[family]
	if f == nil {
		panic("No such family: " + displayName(family, name))
	}
	m := f.metrics[name]
	if m == nil {
		panic("No such metric: " + displayName(family, name))
	}
	m.Put(value)
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
	met.Put(value)
}
