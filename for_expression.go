package main

import "fmt"

func main() {
	// for expression adalah sebuah expression yang digunakan untuk mengecek apakah suatu kondisi benar atau salah
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World", i)
	}

	// for dengan Data Collection (Array dan Slice)
	names := []string{"Raka", "John", "Jane"}

	for index, value := range names {
		fmt.Println("Index", index, "Value", value)
	}

	// for dengan Data Collection (Map)
	newNamesMap := map[string]string{
		"Raka": "Raka",
		"John": "John",
		"Jane": "Jane",
	}

	for index, value := range newNamesMap {
		fmt.Println("Key", index, "Value", value)
	}
}