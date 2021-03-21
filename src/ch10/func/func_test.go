package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int,int){
	return rand.Intn(10),rand.Intn(20)
}

func TestFn(t *testing.T) {
	a,_ :=returnMultiValues()
	t.Log(a)
}


//方法作为参数传递----------------

func timeSpent(myInner func(op int) int) func(op int) int {
	return func(n int) int {
		start :=time.Now()
		ret :=myInner(n)
		fmt.Println("time spent:",time.Since(start).Seconds())
		return  ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second*1)
	fmt.Println("slowFun")
	return op
}

func TestFn02(t *testing.T) {
	a,_ :=returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}