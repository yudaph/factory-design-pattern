package main

import (
	"factory-design-pattern/person"
	"factory-design-pattern/personpro"
	"factory-design-pattern/persontwo"
)

func main() {
	// interface factory
	joko := person.CreatePerson("Joko", 12)
	joko.Talk()

	anto := person.CreatePerson("Joko", 17)
	anto.Talk()

	angela := person.CreatePerson("Angela", 30)
	angela.Talk()

	// factory generator
	kids := persontwo.NewPersonFactory("I am still kids, didn't have ID", false)
	adult := persontwo.NewPersonFactory("I am adult", true)
	john := kids("John", 5)
	john.Talk()

	george := adult("george", 30)
	george.Talk()

	// prototype factory
	herman := personpro.CreatePerson(19)
	herman.Name = "herman"
	herman.Talk()

	volski := personpro.CreatePerson(16)
	volski.Name = "volski"
	volski.Talk()
}
