package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	return &Person{name: name, age: 33}
}

func main_15() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(newPerson("Jon"))

	p := newPerson("Ganbold")
	s := p
	s.name = "Ganbo"
	fmt.Println(p.name)
}
