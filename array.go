package main

import "fmt"

func main() {
	// array adalah sebuah tipe data yang digunakan untuk menyimpan data dalam jumlah banyak
	var names [3]string
	names[0] = "Raka"
	names[1] = "Adalah"
	names[2] = "Ganteng"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// array adalah sebuah tipe data yang digunakan untuk menyimpan data dalam jumlah banyak
	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	fmt.Println(names[0:2])

	var newValues = [...]int{
		90,
		80,
		70,
	}

	fmt.Println(newValues)
	fmt.Println(newValues[0])
	fmt.Println(newValues[1])
	fmt.Println(newValues[2])
	
}