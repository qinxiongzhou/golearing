package empty_interface

import (
	"fmt"
	"testing"
)

/**
空接口与断言
1、空接口可以表示任何类型
2、通过断言来将空接口转换为指定类型
v,ok:=p.(int)
 */

func DoSomething(p interface{})  {
	if i,ok:=p.(int);ok {
		fmt.Println("Integer",i)
		return
	}
	if i,ok:=p.(string);ok {
		fmt.Println("string",i)
		return
	}
	fmt.Println("Unknow Type")
/*
	switch v:=p.(type) {
	case int:
		fmt.Println("Integer",v)
	case string:
		fmt.Println("string",v)
	default:
		fmt.Println("Unknow Type")
	}*/
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
