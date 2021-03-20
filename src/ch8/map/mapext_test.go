package _map

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	//map value is a function
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](2), m[3](2))

}

func TestMapForSet(t *testing.T) {
	mySet := map[string]bool{}
	mySet["平安"]=true
	key := "平安"

	if mySet[key]{
		t.Log(key+" is exist")
	}else {
		t.Log(key + " is not exist")
	}
}