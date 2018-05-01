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

func (customer *Customer) setEmail(email string) {
	customer.email = email
}

func (customer *Customer) getEmail() string {
	return customer.email
}

func main() {
	p3 := newCustomer("Aurelio", "abmf", "124")

	fmt.Println(*p3)

	setName(p3, "nome agora")

	p3.setEmail("abmf@ic.ufal.br")

	fmt.Println("email:", p3.getEmail())
	
}


