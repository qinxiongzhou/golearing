package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employeee struct {
	Id string
	Name string
	Age int
}
//第一种定义方式在实例对应方法被调用是，实例的成员会进行值复制
//func (e Employeee) String() string {
//	//Address is c0000745b0
//	//Address is c0000745e0
//	fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
//}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employeee) String() string {
	//Address is c0000745b0
	//Address is c0000745b0
	fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employeee{"0", "Bob", 20}
	fmt.Printf("Address is %x \n",unsafe.Pointer(&e.Name))
	t.Log(e.String())
}