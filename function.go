package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func setIntroduction(name string, age int, address string) {
	
	fmt.Println("Hello", name, "umur saya", age, "dan saya tinggal di", address)
}

func calculate(numberOne int, numberTwo int) (result int) {
	result = numberOne + numberTwo
	return result
}

// Variadic Function
func variadicFunction(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}	

// Slice Parameter
func sliceParameter(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}


func main() {
	sayHello("Raka")

	age := calculate(10, 20)

	setIntroduction("Raka", age, "Jakarta")

	fmt.Println("age =", calculate(10, 20))

	fmt.Println("total =", variadicFunction(10, 20, 30, 40, 50))

	// Slice Parameter dengan Variadic Function
	sliceNumbers := []int{10, 20, 30, 40, 50}
	fmt.Println("total =", variadicFunction(sliceNumbers...))

	// Slice Parameter dengan Function biasa
	fmt.Println("total =", sliceParameter(sliceNumbers))
}
