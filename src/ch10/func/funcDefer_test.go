package _func

import (
	"fmt"
	"testing"
)

func clear()  {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("Start")
	panic("Fatal error") //defer仍会执行
}
