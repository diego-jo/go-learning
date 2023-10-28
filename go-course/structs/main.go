package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type fck []string

func main() {
	// var person1 person
	// person1.firstName = "diego"
	// person1.lastName = "ol"
	// person1.contact.email = "sdfsdf"
	// person1.contact.phone = "4744444"

	person1 := person{
		firstName: "diego",
		lastName:  "ol",
		contact: contactInfo{
			email: "sdfsdf",
			phone: "ljshdkjhsdfg",
		},
	}

	person1.updateFirstname("jose")
	person1.print()
}

func (p *person) updateFirstname(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
