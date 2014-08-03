package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"testing"
)

func TestRulesCheck(t *testing.T) {
	name := "foo"
	data := metrics.NewProcessStore()
	rule := &Rule{"memory", "rss", LT, 64 * 1024 * 1024, 1, 0, nil}

	// no data in the buffer
	result := rule.Check(name, data)
	assert.Equal(t, Undetermined, result)

	data = metrics.NewProcessStore("memory", "rss", 63*1024*1024)
	result = rule.Check(name, data)
	assert.Equal(t, Triggered, result)
}
