package person

import (
	"fmt"
	"math/rand"
)

type human interface {
	Talk()
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}

// infant
type infant struct {
	person
}

func newInfant(person person) *infant {
	return &infant{person: person}
}

func (i infant) Talk() {
	fmt.Println("mamama")
}

// kid
type kid struct {
	person
}

func newKid(person person) *kid {
	return &kid{person: person}
}

func (k kid) Talk() {
	fmt.Println("Hi my name", k.name, ", my age ", k.age, ", I'm still kid")
}

// teenager
type teenager struct {
	person
	idNumber int
}

func newTeenager(person person, idNumber int) *teenager {
	return &teenager{person: person, idNumber: idNumber}
}

func newTeenagerWithId(person person) *teenager {
	return newTeenager(person, rand.Int())
}

func newTeenagerWithoutId(person person) *teenager {
	return newTeenager(person, 0)
}

func (t teenager) Talk() {
	tellId := ""
	if t.idNumber > 0 {
		tellId = fmt.Sprintf("my id is %d", t.idNumber)
	}
	fmt.Println("Hi my name", t.name, ", my age ", t.age, ", I'm teenager", tellId)
}

// adult
type adult struct {
	person
	idNumber int
}

func newAdult(person person) *adult {
	return &adult{person: person, idNumber: rand.Int()}
}

func (a adult) Talk() {
	fmt.Println("Hi, I'm", a.name, "I am already grown up", "my id", a.idNumber)
}

// CreatePerson was a Person factory
func CreatePerson(name string, age int) human {
	person := *newPerson(name, age)
	if age < 2 {
		return newInfant(person)
	}
	if age < 10 {
		return newKid(person)
	}
	if age < 17 {
		return newTeenagerWithoutId(person)
	}
	if age < 20 {
		return newTeenagerWithId(person)
	}
	return newAdult(person)
}
