package inspeqtor

import (
	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
)

type mockCheckable struct {
	name  string
	store metrics.Store
}

func (c *mockCheckable) Name() string                        { return c.name }
func (c *mockCheckable) Parameter(string) string             { return "" }
func (c *mockCheckable) Metrics() metrics.Store              { return c.store }
func (c *mockCheckable) Resolve([]services.InitSystem) error { return nil }
func (c *mockCheckable) Rules() []*Rule                      { return nil }
func (c *mockCheckable) Verify() []*Event                    { return nil }
func (c *mockCheckable) Collect(bool, func(Checkable))       {}

func MockCheckable(name string) Checkable {
	return &mockCheckable{name, metrics.NewProcessStore("/", 15)}
}
