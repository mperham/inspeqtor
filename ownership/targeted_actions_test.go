package ownership

import (
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/stretchr/testify/assert"
)

func TestBuildAction(t *testing.T) {
	t.Parallel()

	i, err := inspeqtor.New("test", "")
	assert.Nil(t, err)
	assert.NotNil(t, i)

	err = i.Parse()
	assert.Nil(t, err)
	for _, r := range i.Services[0].Rules() {
		for _, a := range r.Actions {
			assert.NotNil(t, a)
		}
	}
}
