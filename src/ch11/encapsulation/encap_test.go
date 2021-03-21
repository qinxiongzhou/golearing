package encapsulation

import "testing"

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{
		Name: "Mike",
		Age:  30,
	}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T",e) //e is encapsulation.Employee
	t.Logf("e2 is %T",e2) //e2 is *encapsulation.Employee
}