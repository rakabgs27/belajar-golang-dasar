# Belajar Golang Dasar

Repositori ini berisi kumpulan materi belajar pemrograman **Go (Golang)** dari dasar. Setiap topik ditulis dalam file `.go` terpisah agar lebih mudah dipahami.

---

## Prasyarat

- [Go](https://golang.org/dl/) versi `1.26.2` atau lebih baru
- Text editor / IDE (VS Code, GoLand, dll)

---

## Cara Menjalankan

Karena setiap file memiliki `func main()` sendiri, jalankan per file menggunakan perintah berikut:

```bash
go run <nama-file>.go
```

Contoh:

```bash
go run helloworld.go
```

---

## Daftar Materi

### 1. Hello World — `helloworld.go`

Program pertama dalam Go. Menampilkan teks ke layar menggunakan `fmt.Println`.

```go
fmt.Println("Hello, World!")
```

---

### 2. Tipe Data Number — `number.go`

Mengenal tipe data angka dan cara mencetaknya.

```go
fmt.Println("Satu =", 1)
fmt.Println("Sepuluh =", 10)
```

---

### 3. Tipe Data Boolean — `boolean.go`

Mengenal nilai `true` dan `false` dalam Go.

```go
fmt.Println("Apakah benar?", true)
fmt.Println("Apakah salah?", false)
```

---

### 4. Tipe Data String — `string.go`

Mengenal string, panjang string dengan `len()`, dan akses karakter berdasarkan index.

```go
fmt.Println(len("Raka"))        // panjang string
fmt.Println(string("Raka"[0]))  // karakter pertama: R
```

---

### 5. Variable — `variable.go`

Cara mendeklarasikan variabel di Go dengan berbagai cara:

| Cara | Contoh |
|------|--------|
| Deklarasi eksplisit | `var name string` |
| Deklarasi dengan nilai | `var lastName = "Programmer"` |
| Shorthand (`:=`) | `numberThree := 30` |
| Multiple variable | `var (cars1 = "Toyota"; cars2 = "BMW")` |

```go
var name string = "Raka"
lastName := "adalah Seorang Programmer"
fmt.Println(name + " " + lastName)
```

---

### 6. Constant — `constant.go`

Konstanta adalah nilai yang tidak bisa diubah setelah dideklarasikan. Menggunakan keyword `const`.

```go
const firstName string = "Raka"

// Multiple constant
const (
    NamaDepan    = "Raka"
    NamaBelakang = "Adalah Seorang Programmer"
)
```

> Mencoba mengubah nilai konstanta akan menyebabkan **error** saat kompilasi.

---

### 7. Boolean (Operasi Perbandingan) — `program_boolean.go`

Menggunakan operator perbandingan (`>=`, `<=`, `==`, dll) yang menghasilkan nilai `bool`.

```go
var nilaiLulus = 80
var lulus = nilaiLulus >= 75  // true
var akhir = nilaiLulus >= 90  // false

fmt.Println("lulus =", lulus)
fmt.Println("akhir =", akhir)
```

---

### 8. Konversi Tipe Data — `conversion.go`

Mengubah tipe data dari satu tipe ke tipe lain secara eksplisit menggunakan casting.

```go
var nilai32 int32 = 32768
var nilai64 int64 = int64(nilai32)  // int32 → int64
var nilai16 int16 = int16(nilai32)  // int32 → int16 (bisa overflow!)
```

> Perhatian: konversi ke tipe yang lebih kecil bisa menyebabkan **overflow**.

---

### 9. Type Declaration — `type_declaration.go`

Membuat tipe data baru berdasarkan tipe yang sudah ada menggunakan keyword `type`.

```go
type noKTP string

var noKTPRaka noKTP = "1234567890"
var noKTPDummy = noKTP("0987654321ABC")
```

> Berguna untuk membuat tipe data yang lebih **spesifik** dan **bermakna** dalam konteks bisnis.

---

### 10. Array — `array.go`

Array adalah kumpulan data dengan tipe yang sama dan ukuran **tetap (fixed)**.

```go
// Deklarasi array dengan ukuran tetap
var names [3]string
names[0] = "Raka"

// Deklarasi array dengan nilai langsung
var values = [3]int{90, 80, 70}

// Ukuran otomatis menggunakan [...]
var newValues = [...]int{90, 80, 70}

fmt.Println(len(names))     // panjang array
fmt.Println(names[0:2])     // slicing array
```

---

### 11. Slice — `slice.go`

Slice mirip dengan array tetapi ukurannya **dinamis (flexible)**. Slice adalah referensi ke array yang mendasarinya.

```go
var fruits = []string{"apple", "banana", "cherry"}

// Slicing
slice1 := fruits[0:2]  // index 0 sampai 1
slice2 := fruits[:2]   // dari awal sampai index 1
slice3 := fruits[1:]   // dari index 1 sampai akhir
slice4 := fruits[:]    // semua elemen
```

**Fungsi penting pada Slice:**

| Fungsi | Kegunaan |
|--------|----------|
| `len(slice)` | Panjang slice (jumlah elemen) |
| `cap(slice)` | Kapasitas slice |
| `append(slice, val)` | Menambahkan elemen ke slice |
| `make([]T, len, cap)` | Membuat slice baru dengan panjang dan kapasitas tertentu |
| `copy(dst, src)` | Menyalin isi slice ke slice lain |

```go
// Membuat slice dengan make
newSlice := make([]string, 2, 5)

// Menambah elemen
newSlice = append(newSlice, "orange")

// Menyalin slice
copySlice := make([]string, len(newSlice), cap(newSlice))
copy(copySlice, newSlice)
```

> **Catatan:** Mengubah elemen pada slice akan **mempengaruhi array asalnya** karena slice adalah referensi. Namun setelah `append` melebihi kapasitas, slice akan membuat backing array baru.

> Di real project, **slice lebih sering digunakan** daripada array karena lebih fleksibel dan efisien.

---

## Struktur File

```
belajar-golang-dasar/
├── helloworld.go        # Hello World pertama
├── number.go            # Tipe data angka
├── boolean.go           # Tipe data boolean
├── string.go            # Tipe data string
├── variable.go          # Deklarasi variabel
├── constant.go          # Konstanta
├── program_boolean.go   # Operasi perbandingan boolean
├── conversion.go        # Konversi tipe data
├── type_declaration.go  # Deklarasi tipe data baru
├── array.go             # Array
├── slice.go             # Slice
├── go.mod               # Go module file
└── README.md            # Dokumentasi ini
```

---

## Urutan Belajar yang Disarankan

1. `helloworld.go` — Mulai dari sini
2. `number.go` — Tipe data angka
3. `boolean.go` — Tipe data boolean
4. `string.go` — Tipe data string
5. `variable.go` — Variabel
6. `constant.go` — Konstanta
7. `program_boolean.go` — Operasi perbandingan
8. `conversion.go` — Konversi tipe data
9. `type_declaration.go` — Type declaration
10. `array.go` — Array
11. `slice.go` — Slice

---

## Referensi

- [Dokumentasi Resmi Go](https://go.dev/doc/)
- [Go Tour (Interaktif)](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
