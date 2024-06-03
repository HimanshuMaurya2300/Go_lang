package main

import "fmt"

type Person struct {
	name   string
	age    int
	height float64
	weight float64
}

type Contact struct {
	email string
	phone string
	fax   string
}

type Address struct {
	Area  string
	city  string
	state string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
}

func main() {

	var person Person
	fmt.Println(person)

	person.name = "Himanshu"
	person.age = 24
	person.height = 5.8
	person.weight = 60

	fmt.Println(person)

	var person1 = Person{"Maurya", 34, 6.5, 80}
	fmt.Println(person1)

	var person2 = new(Person)
	person2.name = "Charlie"
	person2.height = 2.8
	person2.weight = 20

	fmt.Println(person2)
	fmt.Println(person2.name)

	fmt.Println("Age of Himanshu is :", person.age)
	fmt.Println("Age of Charlie is :", person2.age)

	var employee1 Employee

	employee1.Person_Details.name = "Himanshu"
	employee1.Person_Details.age = 24
	employee1.Person_Details.height = 5.8
	employee1.Person_Details.weight = 60
	employee1.Contact_Details.email = "himanshu@gmail.com"
	employee1.Contact_Details.phone = "8765678954"
	employee1.Contact_Details.fax = "fax@123"
	employee1.Address_Details.Area = "Jaitpur"
	employee1.Address_Details.city = "New-Delhi"
	employee1.Address_Details.state = "Delhi"

	fmt.Println(employee1)
	fmt.Println(employee1.Contact_Details)
}
