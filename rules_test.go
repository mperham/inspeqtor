package inspeqtor

import (
	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/metrics/daemon"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	MB = 1024 * 1024
)

func TestRulesCheck(t *testing.T) {
	t.Parallel()
	svc := Service{&Entity{"mysql", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, nil, nil}
	rule := &Rule{&svc, "memory", "rss", LT, "64m", 64 * MB, 0, false, 2, 0, Ok, nil}

	// no data in the buffer
	result := rule.Check(15)
	assert.Equal(t, Ok, rule.State)
	assert.Nil(t, result)

	// Walk thru a series of cycles to verify state transitions
	svc.metrics = metrics.NewProcessStore("/proc", 15)
	loadValue(svc.metrics, "memory", "rss", 65*MB)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 65*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(svc.metrics, "memory", "rss", 63*MB)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 1, rule.TrippedCount)
	assert.Equal(t, 63*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(svc.metrics, "memory", "rss", 62*MB)
	result = rule.Check(15)
	assert.NotNil(t, result)
	assert.Equal(t, result.Type, RuleFailed)
	assert.Equal(t, 2, rule.TrippedCount)
	assert.Equal(t, 62*MB, rule.CurrentValue)
	assert.Equal(t, Triggered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 62*MB)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 3, rule.TrippedCount)
	assert.Equal(t, 62*MB, rule.CurrentValue)
	assert.Equal(t, Triggered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 65*MB)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 65*MB, rule.CurrentValue)
	assert.Equal(t, Recovered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 66*MB)
	result = rule.Check(15)
	assert.NotNil(t, result)
	assert.Equal(t, result.Type, RuleRecovered)
	assert.Equal(t, 66*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)
}

func TestPerSecRulesCheck(t *testing.T) {
	t.Parallel()

	basic := metrics.NewProcessStore("/proc", 15)

	funk := daemon.Sources["mysql"]
	source, err := funk(map[string]string{})
	assert.Nil(t, err)
	assert.NotNil(t, source)
	store := daemon.NewStore(basic, source)
	store.Watch("Queries")
	err = daemon.Prepare(store)
	assert.Nil(t, err)

	svc := Service{&Entity{"mysql", nil, store, nil}, nil, nil, nil}
	rule := &Rule{&svc, "mysql", "Queries", GT, "1k/sec", 1024, 0, true, 2, 0, Ok, nil}

	// no data in the buffer
	result := rule.Check(15)
	assert.Equal(t, Ok, rule.State)
	assert.Nil(t, result)

	// Walk thru a series of cycles to verify state transitions
	loadValue(store, "mysql", "Queries", 1000)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 0, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(store, "mysql", "Queries", 4000)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 3000, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(store, "mysql", "Queries", 20000)
	result = rule.Check(15)
	assert.Nil(t, result)
	assert.Equal(t, 1, rule.TrippedCount)
	assert.Equal(t, 16000, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)
}

func loadValue(store metrics.Store, values ...interface{}) {
	store.(metrics.Loadable).Load(values...)
}
