package main

import "fmt"

type Address struct {
	City string
	Province string
	Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Jakarta", "DKI Jakarta", ""}
	changeCountryToIndonesia(&address)
	fmt.Println(address)
}