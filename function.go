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

// Function as Value
func getHello(name string) string {
	return "Hello " + name
}

// Function as Parameter
func sayHelloWithFunction(name string, function func(string) string) {
	fmt.Println("Hello", function(name))
}

// Function Type Declaration
type Filter func(string) string

func filter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}


// Anonymous Function
type Blacklist func(string) bool

func isBlacklist(name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	}else{
		fmt.Println("You are not blocked", name)
	}
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

	// Function as Value
	hello := getHello
	fmt.Println(hello("Raka"))

	// Function as Parameter
	sayHelloWithFunction("Bagas", getHello)

	// Function Type Declaration
	filter("Raka", getHello)

	// Anonymous Function
	isBlacklist("Raka", func(name string) bool {
		return name == "Bagas"
	})



}
