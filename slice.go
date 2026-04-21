package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "cherry"}
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))
	fmt.Println(fruits[0:2])
	fmt.Println(fruits[0:1])

	// slice 1 adalah sebuah slice yang berisi data dari index 0 sampai 2
	slice1 := fruits[0:2]
	fmt.Println(slice1)

	// slice 2 adalah sebuah slice yang berisi data dari index 0 sampai 2
	slice2 := fruits[:2]
	fmt.Println(slice2)

	// slice 3 adalah sebuah slice yang berisi data dari index 1 sampai akhir
	slice3 := fruits[1:]
	fmt.Println(slice3)

	// slice 4 adalah sebuah slice yang berisi data dari index 0 sampai akhir
	slice4 := fruits[:]
	fmt.Println(slice4)

	slice4[0] = "appleNew"

	fmt.Println(fruits)

	// Menggunakan len untuk mengetahui panjang slice
	fmt.Println(len(slice4))

	// Menggunakan cap untuk mengetahui kapasitas slice
	fmt.Println(cap(slice3))

	// Menggunakan append untuk menambahkan data ke slice
	slice4 = append(slice4, "orange")
	fmt.Println(slice4)
	fmt.Println(fruits)

	// Menggunakan make untuk membuat slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "apple"
	newSlice[1] = "banana"
	
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	
	newSlice2 := append(newSlice, "orange")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// menggunakan copy untuk menyalin data dari slice ke slice lain
	copySlice := make([]string, len(newSlice2), cap(newSlice2))
	copy(copySlice, newSlice2)
	fmt.Println(newSlice2)
	fmt.Println(copySlice)

	// Di real project lebih banyak menggunakan slice daripada array
	// karena slice lebih fleksibel dan mudah untuk digunakan
	// dan slice lebih efisien dalam hal penggunaan memori
	// dan slice lebih efisien dalam hal penggunaan waktu
	// dan slice lebih efisien dalam hal penggunaan prosesor
	// dan slice lebih efisien dalam hal penggunaan baterai
	// dan slice lebih efisien dalam hal penggunaan internet
	// dan slice lebih efisien dalam hal penggunaan data
}
	
	