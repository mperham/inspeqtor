package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"testing"
)

const (
	MB = 1024 * 1024
)

func TestRulesCheck(t *testing.T) {
	t.Parallel()
	svc := Service{&Entity{"mysql", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, nil, nil}
	rule := &Rule{&svc, "memory", "rss", LT, "64m", 64 * MB, 0, 2, 0, Ok, nil}

	// no data in the buffer
	result := rule.Check()
	assert.Equal(t, Ok, rule.State)
	assert.Nil(t, result)

	// Walk thru a series of cycles to verify state transitions
	svc.metrics = metrics.NewProcessStore("/proc", 15)
	loadValue(svc.metrics, "memory", "rss", 65*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 65*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(svc.metrics, "memory", "rss", 63*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 1, rule.TrippedCount)
	assert.Equal(t, 63*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)

	loadValue(svc.metrics, "memory", "rss", 62*MB)
	result = rule.Check()
	assert.NotNil(t, result)
	assert.Equal(t, result.Type, RuleFailed)
	assert.Equal(t, 2, rule.TrippedCount)
	assert.Equal(t, 62*MB, rule.CurrentValue)
	assert.Equal(t, Triggered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 62*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 3, rule.TrippedCount)
	assert.Equal(t, 62*MB, rule.CurrentValue)
	assert.Equal(t, Triggered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 65*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 65*MB, rule.CurrentValue)
	assert.Equal(t, Recovered, rule.State)

	loadValue(svc.metrics, "memory", "rss", 66*MB)
	result = rule.Check()
	assert.NotNil(t, result)
	assert.Equal(t, result.Type, RuleRecovered)
	assert.Equal(t, 66*MB, rule.CurrentValue)
	assert.Equal(t, Ok, rule.State)
}

func loadValue(store metrics.Store, values ...interface{}) {
	store.(metrics.Loadable).Load(values...)
}
