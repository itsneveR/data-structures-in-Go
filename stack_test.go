package ds

import "testing"

type testStruct struct {
	arg, exV1, exV2 int
}

func TestNew(t *testing.T) {

	stack := New(5)

	t.Log(stack)
	if stack.top != -1 {
		t.Errorf("nope")
	}

}
