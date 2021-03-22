package polymorphism

import "testing"

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer)WriteHelloWorld() Code {
	return "fmt.Println(\"Go Programmer say : Hello World!\")"
}

type JavaProgrammer struct {

}

func (java * JavaProgrammer)WriteHelloWorld() Code {
	return "fmt.Println(\"Java Programmer say : Hello World!\")"
}

func TestPolymorphism(t *testing.T) {
	var p1 Programmer = new(GoProgrammer)
	var p2 Programmer = new(JavaProgrammer)
	p1.WriteHelloWorld()
	p2.WriteHelloWorld()
}
