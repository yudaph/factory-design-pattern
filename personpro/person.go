package personpro

import (
	"fmt"
	"math/rand"
)

type person struct {
	Name, talk    string
	age, idNumber int
}

func (p person) Talk() {
	fmt.Println(p.talk, p.Name)
}

func newPerson(talk string, age, idNumber int) *person {
	return &person{
		talk:     talk,
		age:      age,
		idNumber: idNumber,
	}
}

func CreatePerson(age int) *person {
	if age < 2 {
		return newPerson("mamama", age, 0)
	}
	if age < 10 {
		return newPerson("I am still kids", age, 0)
	}
	if age < 17 {
		return newPerson("I am teenager, but didn't have ID", age, 0)
	}
	if age < 20 {
		id := rand.Int()
		return newPerson(fmt.Sprintf("I am teenager, this is my id %d", id), age, id)
	}
	id := rand.Int()
	return newPerson(fmt.Sprintf("I am adult, this is my id %d", id), age, id)
}
