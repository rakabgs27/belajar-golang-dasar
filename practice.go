package main

import "fmt"

type User struct {
    name string
    age  int
}

func updateAge(u *User, newAge int) {
    u.age = newAge
}

func main() {
    user := User{name: "Alice", age: 25}
    fmt.Println("Before:", user) // Before : name : "Alice", age: "25"

    updateAge(&user, 30)
    fmt.Println("After: ", user) // After : name : "Alice", age: "30"
}