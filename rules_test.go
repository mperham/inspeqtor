package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"testing"
)

func TestRulesCheck(t *testing.T) {
	name := "foo"
	data := metrics.NewStore()
	rule := &Rule{"memory", "rss", LT, 64 * 1024 * 1024, 1, 0, nil}

	// no data in the buffer
	result := rule.Check(name, data)
	assert.Equal(t, Undetermined, result)

	data = metrics.NewStore("memory", "rss", 63*1024*1024)
	result = rule.Check(name, data)
	assert.Equal(t, Triggered, result)
}
