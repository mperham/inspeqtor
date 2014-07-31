package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"testing"
)

func TestRulesCheck(t *testing.T) {
	name := "foo"
	data := metrics.NewStore()
	rule := &Rule{"memory", "rss", GT, 64 * 1024 * 1024, 1, Ok, nil}

	// no data in the buffer
	result := checkRule(name, data, rule)
	assert.Equal(t, Unchanged, result)

	data = metrics.NewStore("memory", "rss", 65*1024*1024)
	result = checkRule(name, data, rule)
	assert.Equal(t, Unchanged, result)
}
