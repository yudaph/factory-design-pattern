package main

import (
	"factory-design-pattern/person"
	"factory-design-pattern/personpro"
	"factory-design-pattern/persontwo"
)

func main() {
	// interface factory
	joko := person.CreatePerson("Joko", 12)
	joko.Talk() // Hi my name Joko , my age  12 , I'm teenager

	anto := person.CreatePerson("Joko", 17)
	anto.Talk() // Hi my name Joko , my age  17 , I'm teenager my id is 5577006791947779410

	angela := person.CreatePerson("Angela", 30)
	angela.Talk() // Hi, I'm Angela I am already grown up my id 8674665223082153551  

	// factory generator
	kids := persontwo.NewPersonFactory("I am still kids, didn't have ID", false)
	adult := persontwo.NewPersonFactory("I am adult", true)
	john := kids("John", 5)
	john.Talk() // Hi, my name John age 5 I am still kids, didn't have ID  

	george := adult("george", 30)
	george.Talk() // Hi, my name george age 30 I am adult this is my id 6129484611666145821 

	// prototype factory
	herman := personpro.CreatePerson(19)
	herman.Name = "herman"
	herman.Talk() // I am teenager, this is my id 4037200794235010051 herman 

	volski := personpro.CreatePerson(16)
	volski.Name = "volski"
	volski.Talk() // I am teenager, but didn't have ID volski
}
