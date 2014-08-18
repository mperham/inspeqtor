package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"testing"
)

const (
	MB = 1024 * 1024
)

func TestRulesCheck(t *testing.T) {
	t.Parallel()
	svc := Service{"mysql", 0, services.Down, nil, nil, metrics.NewProcessStore(), nil}
	rule := &Rule{&svc, "memory", "rss", LT, "64m", 64 * MB, 0, 2, 0, Ok, nil}

	// no data in the buffer
	result := rule.Check()
	assert.Equal(t, Ok, rule.state)
	assert.Nil(t, result)

	// Walk thru a series of cycles to verify state transitions
	svc.Metrics = metrics.NewProcessStore("memory", "rss", 65*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 65*MB, rule.currentValue)
	assert.Equal(t, Ok, rule.state)

	svc.Metrics = metrics.NewProcessStore("memory", "rss", 63*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 1, rule.trippedCount)
	assert.Equal(t, 63*MB, rule.currentValue)
	assert.Equal(t, Ok, rule.state)

	svc.Metrics = metrics.NewProcessStore("memory", "rss", 62*MB)
	result = rule.Check()
	assert.NotNil(t, result)
	assert.Equal(t, 2, rule.trippedCount)
	assert.Equal(t, 62*MB, rule.currentValue)
	assert.Equal(t, Triggered, rule.state)

	svc.Metrics = metrics.NewProcessStore("memory", "rss", 62*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 3, rule.trippedCount)
	assert.Equal(t, 62*MB, rule.currentValue)
	assert.Equal(t, Triggered, rule.state)

	svc.Metrics = metrics.NewProcessStore("memory", "rss", 65*MB)
	result = rule.Check()
	assert.NotNil(t, result)
	assert.Equal(t, 65*MB, rule.currentValue)
	assert.Equal(t, Recovered, rule.state)

	svc.Metrics = metrics.NewProcessStore("memory", "rss", 66*MB)
	result = rule.Check()
	assert.Nil(t, result)
	assert.Equal(t, 66*MB, rule.currentValue)
	assert.Equal(t, Ok, rule.state)
}
