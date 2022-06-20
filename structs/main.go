package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 123456,
		},
	}
	alexPointer := &alex
	(*alexPointer).updateFirstName("blex")
	alex.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (personPointer *person) updateFirstName(firstName string) {
	personPointer.firstName = firstName
}
