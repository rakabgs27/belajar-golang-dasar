package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	} else {
		fmt.Println("App Running")
	}
}
func main() {
	runApp(true)
}