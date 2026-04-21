package main

import "fmt"

func main() {
	const firstName string = "Raka"
	const lastName string = "Adalah Seorang Programmer"

	// error karena constant tidak bisa diubah
	// firstName = "John"
	//lastName = "Doe"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// multiple constant
	const (
		NamaDepan = "Raka"
		NamaBelakang = "Adalah Seorang Programmer"
	)

	fmt.Println(NamaDepan)
	fmt.Println(NamaBelakang)
}