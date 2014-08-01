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
	data map[string]map[string]*util.RingBuffer
}

func NewStore(values ...interface{}) Storage {
	s := Storage{
		make(map[string]map[string]*util.RingBuffer),
	}
	if len(values) > 0 {
		fam := values[0].(string)
		name := values[1].(string)
		for _, val := range values[2:] {
			s.save(fam, name, int64(val.(int)))
		}
	}
	return s
}

func (store Storage) save(family string, name string, value int64) {
	fam := store.data[family]
	if fam == nil {
		store.data[family] = map[string]*util.RingBuffer{}
		fam = store.data[family]
	}

	data := fam[name]
	if data == nil {
		fam[name] = util.NewRingBuffer(SLOTS)
		data = fam[name]
	}

	data.Add(value)
}

func Lookup(store Storage, family string, name string) int64 {
	return store.Get(family, name)
}

func LookupAt(store Storage, family string, name string, idx int) int64 {
	return store.GetAt(family, name, idx)
}

func (store Storage) GetAt(family string, name string, idx int) int64 {
	buf := store.data[family][name]
	if buf == nil {
		return -1
	}
	return buf.At(-1 * idx).(int64)
}

func (store Storage) Get(family string, name string) int64 {
	return store.GetAt(family, name, 0)
}
