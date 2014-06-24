package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "regexp"
  "strconv"
  "strings"
)

type SystemMetrics struct {
  When float32
  UserCpu float32
  SystemCpu float32
  FreeMem int
  Load1 float32
  Load5 float32
  Load15 float32
  PercentSwapInUse int8
}

var (
  meminfoParser *regexp.Regexp = regexp.MustCompile("([^:]):\\s+(\\d+)")
)

func CollectSystemMetrics() (*SystemMetrics, error) {

  var err error
  var metrics *SystemMetrics = &SystemMetrics{}

  err = collectLoadAverage(metrics)
  if err != nil { return nil, err }
  err = collectMemory(metrics)
  if err != nil { return nil, err }

  return metrics, nil
}

func collectMemory(metrics *SystemMetrics) error {
  contentBytes, err := ioutil.ReadFile("/proc/meminfo")
  if err != nil { return err }
  content := bytes.NewBuffer(contentBytes).String()
  lines := strings.Split(content, "\n")

  var memMetrics map[string]int
  for _, line := range(lines) {
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
  if free == 0 {
    metrics.PercentSwapInUse = 100
  } else if free == memMetrics["SwapTotal"] {
    metrics.PercentSwapInUse = 100
  } else {
    metrics.PercentSwapInUse = 100 - int8(100*(memMetrics["SwapFree"] / memMetrics["SwapTotal"]))
  }

  return nil
}

func collectLoadAverage(metrics *SystemMetrics) error {
  contentBytes, err := ioutil.ReadFile("/proc/loadavg")
  if err != nil { return err }

  content := bytes.NewBuffer(contentBytes).String()
  slices := strings.Split(content, " ")
  load1, err := strconv.ParseFloat(slices[0], 32)
  if err != nil { return err }
  load5, err := strconv.ParseFloat(slices[1], 32)
  if err != nil { return err }
  load15, err := strconv.ParseFloat(slices[2], 32)
  if err != nil { return err }


  metrics.Load1 = float32(load1)
  metrics.Load5 = float32(load5)
  metrics.Load15 = float32(load15)
  return nil
}
