package main

import "fmt"

func cekStatusGamer(name string) (int, bool, string) {
    // return skor, isVIP, level
    return 100, true, "Grandmaster"
}

func main() {
	// if expression adalah sebuah expression yang digunakan untuk mengecek apakah suatu kondisi benar atau salah
	var name = "Raka Bagas"

	if name == "Raka" {
		fmt.Println("Hello Raka")
	} else {
		fmt.Println("Hello World")
	}

	// if expression dengan multiple statement
	var age = 15

	if age > 18 {
		fmt.Println("Anda sudah dewasa")
	} else if age > 12 {
		fmt.Println("Anda masih remaja")
	} else {
		fmt.Println("Anda masih anak-anak")
	}

	// if expression dengan short statement
	var newName = "Raka Bagas"

	if length := len(newName); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar", newName)
	}

	// Kita deklarasikan 3 variabel sekaligus di short statement
    if score, vip, level := cekStatusGamer("Raka"); vip {
        fmt.Printf("User %s (Skor: %d) sedang online!", level, score)
    }
}