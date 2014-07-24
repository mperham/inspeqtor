package metrics

import (
	"time"
)

type ProcessMetrics struct {
	When           time.Time
	UserCpu        uint64
	SystemCpu      uint64
	UserChildCpu   uint64
	SystemChildCpu uint64
	RSS            uint64
	VMSize         uint64
}

type Process struct {
	MetricsHistory []*ProcessMetrics
}

func (p *Process) CaptureMetrics() (*ProcessMetrics, error) {
	return &ProcessMetrics{time.Now(), 0, 0, 0, 0, 0, 0}, nil
}
