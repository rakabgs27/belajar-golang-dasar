package main

import (
	"fmt"
	"belajar-golang-dasar/database"
	_"belajar-golang-dasar/internal"
)

func main() {
	fmt.Println(database.GetConnection())
}