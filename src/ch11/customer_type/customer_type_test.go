package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(myInner IntConv) IntConv {
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

	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}