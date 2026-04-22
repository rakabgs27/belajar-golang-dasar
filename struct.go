package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

// struct method
func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {

	var customer Customer
	customer.Name = "Raka"
	customer.Address = "Jakarta"
	customer.Age = 20

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	//struct literal
	customer2 := Customer{
		Name: "Bagas",
		Address: "Bandung",
		Age: 20,
	}
	fmt.Println(customer2)
	fmt.Println(customer2.Name)
	fmt.Println(customer2.Address)
	fmt.Println(customer2.Age)

	customer3 := Customer{"Fitri", "Malang", 20}
	fmt.Println(customer3)

	// struct method
	customer3.sayHello()
}