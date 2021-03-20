package fib

import "testing"

func TestFib(t *testing.T) {
	var a = 1
	var b = 1

	for i := 1; i < 10; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a

	}
}

//交换两个变量
func TestExchange(t *testing.T) {

	a:=1
	b:=2

	//tmp:=a
	//a = b
	//b = tmp

	a , b = b, a

	t.Log(a,b)

}