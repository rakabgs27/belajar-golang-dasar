package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func main() {
	// address1 := Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	// address2 := address1

	// address2.City = "Jakarta Selatan"

	// fmt.Println(address1)
	// fmt.Println(address2)

	var address1 Address = Address{"Jakarta Pusat", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jakarta Selatan"
	fmt.Println(address1)
	fmt.Println(address2)
}