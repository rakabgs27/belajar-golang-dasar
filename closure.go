package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		counter++
		fmt.Println("increment")
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)

	counter2 := 0
	increment2 := func(c int) {  // nilai disalin ke parameter
		c++
		fmt.Println("increment")
	}
	
	increment2(counter2)
	increment2(counter2)
	increment2(counter2)
	fmt.Println(counter2) // tetap 0!


}