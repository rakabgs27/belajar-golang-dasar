package main

import "fmt"

func main() {
	// type declaration adalah sebuah fitur di golang yang memungkinkan kita untuk mendeklarasikan tipe data baru berdasarkan tipe data yang sudah ada
	// type declaration biasanya digunakan untuk membuat tipe data baru yang lebih spesifik dan lebih mudah untuk digunakan
	type noKTP string

	// contoh penggunaan type declaration
	var noKTPRaka noKTP = "1234567890"

	// contoh penggunaan type declaration dengan tipe data yang sudah ada
	var noKTPDummy  = noKTP("0987654321ABC")

	// contoh penggunaan type declaration dengan tipe data string
	var contoh = "ABCDEFG"

	// contoh penggunaan type declaration dengan tipe data string
	var noKTPContoh = noKTP(contoh)

	fmt.Println("noKTPRaka =", noKTPRaka)
	fmt.Println("noKTPDummy =", noKTPDummy)
	fmt.Println("noKTPContoh =", noKTPContoh)
}