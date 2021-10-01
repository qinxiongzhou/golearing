package main

import "testing"

func TestName(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 6}
	t.Log(m1[2])                 //4
	t.Logf("len m1=%d", len(m1)) //len m1=3
	m2 := map[int]int{}
	t.Logf("len m2=%d", len(m2)) //len m2=0
	m3 := make(map[int]int, 10)  //用make创建map结构，10指cap，不是len
	t.Logf("len m3=%d", len(m3)) //len m3=0
}

func TestAccessNotExiststringKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	//value 不存在，返回值也是0，用如下方式区分值为0还是不存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 value is %d", v)
	} else {
		t.Logf("key 3 is not existing")
	}

	m1[3] = 0
	//value 不存在，返回值也是0，用如下方式区分值为0还是不存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3 value is %d", v)
	} else {
		t.Logf("key 3 is not existing")
	}
}

//map中放个map
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 6}
	for k, v := range m1 {
		t.Log(k, v)
	}

	m2 := map[string]map[string]int{"1": {"1": 1, "1.1": 1}, "2": {"2": 2}}
	for k, v := range m2 {
		t.Log(k)
		for subk, subv := range v {
			t.Log(subk, subv)
		}
	}
}

type Person struct {
	age  int
	name string
}

//map中存放对象
func TestStructMap(t *testing.T) {
	class := map[int]*Person{}
	p1 := &Person{
		age:  11,
		name: "dalei",
	}

	addage(p1)
	class[1] = p1

	t.Log(class)
	t.Log(p1)
}

func addage(p *Person) *Person {
	p.age += 1
	return p
}
