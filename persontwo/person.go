package persontwo

import (
	"fmt"
	"math/rand"
)

type person struct {
	name, talk    string
	age, idNumber int
}

func (p person) Talk() {
	if p.idNumber > 0 {
		fmt.Println("Hi, my name", p.name, "age", p.age, p.talk, "this is my id", p.idNumber)
		return
	}
	fmt.Println("Hi, my name", p.name, "age", p.age, p.talk)
}

func NewPersonFactory(talk string, createId bool) func(name string, age int) *person {
	return func(name string, age int) *person {
		idNumber := 0
		if createId {
			idNumber = rand.Int()
		}
		return &person{
			name:     name,
			talk:     talk,
			age:      age,
			idNumber: idNumber,
		}
	}
}
