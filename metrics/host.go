package metrics

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SystemMetrics struct {
	When             time.Time
	CpuUsage         CpuMetrics
	FreeMem          uint64
	Load1            float32
	Load5            float32
	Load15           float32
	PercentSwapInUse int8
}

type CpuMetrics struct {
	Total     uint64
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	Irq       uint64
	SoftIrq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

var (
	meminfoParser *regexp.Regexp = regexp.MustCompile("([^:]+):\\s+(\\d+)")
)

func CollectHostMetrics(path string) (*SystemMetrics, error) {
	var err error
	var metrics *SystemMetrics = &SystemMetrics{}
	metrics.When = time.Now()

	err = collectLoadAverage(path, metrics)
	if err != nil {
		return nil, err
	}
	err = collectMemory(path, metrics)
	if err != nil {
		return nil, err
	}
	err = collectCpu(path, metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func collectMemory(path string, metrics *SystemMetrics) error {
	contentBytes, err := ioutil.ReadFile(path + "/meminfo")
	if err != nil {
		return err
	}
	lines := strings.Split(string(contentBytes), "\n")

	memMetrics := make(map[string]uint64)
	for _, line := range lines {
		if line == "" {
			continue
		}

		results := meminfoParser.FindStringSubmatch(line)
		if results == nil {
			log.Println("Unknown input: " + line)
			continue
		}
		val, err := strconv.ParseUint(results[2], 10, 64)
		if err != nil {
			log.Println("Unexpected input: " + results[2] + " in " + line)
			return err
		}
		memMetrics[results[1]] = uint64(val)
	}

	metrics.FreeMem = memMetrics["MemFree"]
	free := memMetrics["SwapFree"]
	total := memMetrics["SwapTotal"]
	if free == 0 {
		metrics.PercentSwapInUse = 100
	} else if free == total {
		metrics.PercentSwapInUse = 100
	} else {
		metrics.PercentSwapInUse = 100 - int8(100*(float64(free)/float64(total)))
	}

	return nil
}

func collectLoadAverage(path string, metrics *SystemMetrics) error {
	contentBytes, err := ioutil.ReadFile(path + "/loadavg")
	if err != nil {
		return err
	}

	slices := strings.Split(string(contentBytes), " ")
	load1, err := strconv.ParseFloat(slices[0], 32)
	if err != nil {
		return err
	}
	load5, err := strconv.ParseFloat(slices[1], 32)
	if err != nil {
		return err
	}
	load15, err := strconv.ParseFloat(slices[2], 32)
	if err != nil {
		return err
	}

	metrics.Load1 = float32(load1)
	metrics.Load5 = float32(load5)
	metrics.Load15 = float32(load15)
	return nil
}

func collectCpu(path string, metrics *SystemMetrics) error {
	contents, err := ioutil.ReadFile(path + "/stat")
	if err != nil {
		return err
	}

	lines := strings.Split(string(contents), "\n")
	line := lines[0]
	fields := strings.Fields(line)
	metrics.CpuUsage = createCpuMetrics(fields)
	return nil
}

func createCpuMetrics(fields []string) CpuMetrics {
	s := CpuMetrics{}
	s.User, _ = strconv.ParseUint(fields[1], 10, 64)
	s.Nice, _ = strconv.ParseUint(fields[2], 10, 64)
	s.System, _ = strconv.ParseUint(fields[3], 10, 64)
	s.Idle, _ = strconv.ParseUint(fields[4], 10, 64)
	s.IOWait, _ = strconv.ParseUint(fields[5], 10, 64)
	s.Irq, _ = strconv.ParseUint(fields[6], 10, 64)
	s.SoftIrq, _ = strconv.ParseUint(fields[7], 10, 64)
	s.Steal, _ = strconv.ParseUint(fields[8], 10, 64)
	s.Guest, _ = strconv.ParseUint(fields[9], 10, 64)
	s.GuestNice, _ = strconv.ParseUint(fields[10], 10, 64)
	s.Total = s.User + s.Nice + s.System + s.Idle + s.IOWait +
		s.Irq + s.SoftIrq + s.Steal + s.Guest + s.GuestNice
	return s
}
