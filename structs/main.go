package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	alexis := person{
		firstName: "Alexis",
		lastName: "Bellet",
		contactInfo: contactInfo{
			email: "example@gmail.com",
			zipCode: 10001,
		},
	}
	fmt.Printf("%+v", alexis)
	// assign zero-values (i.e. empty string for string types & 0 for int types)
	var james person
	fmt.Printf("%+v", james)
}