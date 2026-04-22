package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func saveData(name string) error {
	if name == "" {
		return &ValidationError{Message: "name is required"}
	} else {
		return nil
	}
}	

func main() {
	err := saveData("")
	if err != nil {
		switch finalError := err.(type) {
			case *ValidationError:
				fmt.Println("Validation Error : ", finalError.Error())
			default:
				fmt.Println("Error : ", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}