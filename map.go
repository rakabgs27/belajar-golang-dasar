package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Raka",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	// menambahkan data ke map
	person["title"] = "Programmer"
	fmt.Println(person)

	// menghapus data dari map
	delete(person, "title")
	fmt.Println(person)
}