package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("Pet Speak")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("Pet SpeakTo ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) SpeakTo(host string) {
	fmt.Println("Dog SpeakTo", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Pet.SpeakTo("Hello World!")
	dog.SpeakTo("Hello World!")
}