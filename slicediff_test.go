package slicediff

import (
	"testing"

	"github.com/remogatto/prettytest"
)

type testSuite struct {
	prettytest.Suite
}

type Foo struct {
	Field string
}

type Bar struct {
	Field string
}

type Bars []Bar
type Foos []Foo

func (f Foos) Strings() (result []string) {
	for _, el := range f {
		result = append(result, el.Field)
	}
	return
}

func (f Bars) Strings() (result []string) {
	for _, el := range f {
		result = append(result, el.Field)
	}
	return
}

func TestRunner(t *testing.T) {
	prettytest.Run(
		t,
		new(testSuite),
	)
}

func (t *testSuite) TestDiff() {
	src := Foos{
		Foo{"A"},
		Foo{"B"},
		Foo{"C"},
	}
	dst := Bars{
		Bar{"A"},
	}

	actions := Diff(dst, src)
	t.True(len(actions) > 0)

	t.Equal(Add, actions["B"].Type)
	t.Equal(1, actions["B"].Id)
}
