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

func main() {
	sayHello("Raka")

	age := calculate(10, 20)

	setIntroduction("Raka", age, "Jakarta")

	fmt.Println("age =", calculate(10, 20))
}
