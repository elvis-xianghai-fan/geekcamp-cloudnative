package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) SetName(newName string) {
	p.Name = newName
}

func (p *Person) SetName2(newName string) {
	p.Name = newName
}

func main() {
	person := new(Person)
	person.Name = "init"

	person.SetName("name 1")
	fmt.Println(person)

	person.SetName2("name 2")
	fmt.Println(person)
}
