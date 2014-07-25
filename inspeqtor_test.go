package inspeqtor

import (
	"testing"
)

func TestInspeqtorParse(t *testing.T) {
	i, err := New("test")
	ok(t, err)
	err = i.Parse()
	ok(t, err)
	equals(t, uint16(15), i.GlobalConfig.Top.CycleTime)
}

func assert(tb testing.TB, condition bool, msg string) {
	if !condition {
		tb.Error(msg)
	}
}
func ok(tb testing.TB, err error) {
	if err != nil {
		tb.Fatal(err)
	}
}
func equals(tb testing.TB, exp, act interface{}) {
	if exp != act {
		tb.Error("Expected", exp, ", received ", act)
	}
}
