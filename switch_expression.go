package main

import "fmt"

func main() {
	var name = "Raka"

	// switch expression adalah sebuah expression yang digunakan untuk mengecek apakah suatu kondisi benar atau salah
	switch name {
		case "Raka":
			fmt.Println("Hello Raka")
		case "John":
			fmt.Println("Hello John")
		default:
			fmt.Println("Hello World")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Nama terlalu panjang")
		case false:
			fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	switch {
		case name == "Raka":
			fmt.Println("Hello Raka")
		case name == "John":
			fmt.Println("Hello John")
		default:
			fmt.Println("Hello World")
	}
}