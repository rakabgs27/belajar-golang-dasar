package main

import "fmt"

func main() {
	type noKTP string

	var noKTPRaka noKTP = "1234567890"

	var noKTPDummy  = noKTP("0987654321ABC")

	var contoh = "ABCDEFG"

	var noKTPContoh = noKTP(contoh)

	fmt.Println("noKTPRaka =", noKTPRaka)
	fmt.Println("noKTPDummy =", noKTPDummy)
	fmt.Println("noKTPContoh =", noKTPContoh)
}