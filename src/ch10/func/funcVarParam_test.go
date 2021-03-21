package _func

import "testing"

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1,2,3,5))
	t.Log(sum(1,5,3,12))
}