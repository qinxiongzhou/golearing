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

func TestUint32(t *testing.T) {
	myMap := make(map[uint32]uint32)
	myMap[111111]=45
	myMap[222222]=53
	myMap[333333]=62

	var sum uint32 = 0

	for _, v := range myMap {
		sum = sum + v
	}
	
	var couponNum uint32 = 20
	
	count1 := myMap[111111] *couponNum/sum
	count2 := myMap[222222] *couponNum/sum
	count3 := couponNum - count1 - count2


	t.Log(count1)
	t.Log(count2)
	t.Log(count3)
	t.Log(count1 + count2 + count3)
}

func TestInt(t *testing.T) {

	myMap := make(map[int]int)
	myMap[111111]=45
	myMap[222222]=53
	myMap[333333]=62

	var sum int = 0

	for _, v := range myMap {
		sum = sum + v
	}

	var couponNum int = 4


	count1 := myMap[111111] *couponNum/sum
	count2 := myMap[222222] *couponNum/sum

	count3 := couponNum - count1 - count2


	t.Log(count1)
	t.Log(count2)
	t.Log(count3)
	t.Log(count1 + count2 + count3)
}