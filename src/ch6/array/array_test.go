package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr3 := [...]int{1,3,4,5} //不关注具体个数，可以用[...]表示
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	for i := 0;i<len(arr3);i++ {
		t.Log(arr3[i])
	}

	// idx：数组下标；e；数组值
	for idx, e := range arr3 {
		t.Log(idx,e)
	}

	//当不关注下标值时，可以用_代替
	for _, i := range arr3 {
		t.Log(i)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	arr3_sec := arr3[1:2]
	t.Log(arr3_sec) //[3]
	arr3_sec = arr3[1:3]
	t.Log(arr3_sec) //[3 4]
	arr3_sec = arr3[1:len(arr3)]
	t.Log(arr3_sec) //[3 4 5]
	arr3_sec = arr3[1:]
	t.Log(arr3_sec) //[3 4 5]
	arr3_sec = arr3[:3]
	t.Log(arr3_sec) //[1 3 4]
}