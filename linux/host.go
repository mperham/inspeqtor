package linux

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type SystemMetrics struct {
	When             float32
	CpuUsage         *CpuMetrics
	FreeMem          int
	Load1            float32
	Load5            float32
	Load15           float32
	PercentSwapInUse int8
}

type CpuMetrics struct {
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

func CollectSystemMetrics(path string) (*SystemMetrics, error) {

	var err error
	var metrics *SystemMetrics = &SystemMetrics{}

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

	memMetrics := make(map[string]int)
	for _, line := range lines {
		results := meminfoParser.FindStringSubmatch(line)
		if results == nil {
			fmt.Println("Unknown input: " + line)
			continue
		}
		val, err := strconv.Atoi(results[2])
		if err != nil {
			fmt.Println("Unexpected input: " + results[2] + " in " + line)
			return err
		}
		memMetrics[results[1]] = val
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

func createCpuMetrics(fields []string) *CpuMetrics {
	s := &CpuMetrics{}
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
	return s
}
