package main 

import (
	"fmt"
)

type Customer struct {
	name string
	email string
	password string
}

func newCustomer(name, email, password string) *Customer {
	return &Customer{
		name: name,
		email: email,
		password: password}
}	

func setName(customer *Customer, name string) {
	customer.name = name
}

func main() {
	p3 := newCustomer("Aurelio", "abmf", "124")

	fmt.Println(*p3)

	setName(p3, "nome agora")

	fmt.Println(*p3)
	
}


