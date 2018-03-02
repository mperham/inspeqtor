package statsd

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/metrics"
)

var (
	hostname string
)

func init() {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hostname = h
}

func Dial(address string) (net.Conn, error) {
	return net.Dial("udp", address)
}

func Export(statsd io.Writer, i *inspeqtor.Inspeqtor) error {
	i.Host.Metrics().Each(func(family, name string, metric metrics.Metric) {
		writeStatsd(statsd, "host", family, name, metric)
	})
	for _, x := range i.Services {
		x.Metrics().Each(func(family, name string, metric metrics.Metric) {
			writeStatsd(statsd, x.Name(), family, name, metric)
		})
	}
	return nil
}

func writeStatsd(statsd io.Writer, thing string, family, name string, metric metrics.Metric) {
	typeCode := "c"
	if metric.Type() == metrics.Gauge {
		typeCode = "g"
	}
	if name == "" {
		statsd.Write([]byte(fmt.Sprintf("%s.%s.%s:%.2f|%s\n", hostname, thing, family, metric.Get(), typeCode)))
	} else {
		statsd.Write([]byte(fmt.Sprintf("%s.%s.%s.%s:%.2f|%s\n", hostname, thing, family, name, metric.Get(), typeCode)))
	}
}
