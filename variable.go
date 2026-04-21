package main

import "fmt"

func main() {
	var name string
	name = "Raka"
	fmt.Println(name)
	var lastName = "adalah Seorang Programmer"
	fmt.Println(lastName)

	fmt.Println(name + " " + lastName)

	var numberOne int8

	numberOne = 10

	fmt.Println(numberOne)

	fmt.Println(numberOne + 10)

	var numberTwo = 10

	fmt.Println("numberTwo =", numberTwo)

	// shorthand variable
	numberThree := 30

	fmt.Println("numberThree =", numberThree)

	// multiple variable
	var (
		cars1 = "Toyota"
		cars2 = "BMW"
	)

	fmt.Println(cars1)
	fmt.Println(cars2)
}