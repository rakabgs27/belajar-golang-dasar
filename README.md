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

### 12. Map — `map.go`

Map adalah struktur data **key-value pair**. Key harus unik, dan value bisa bertipe apa saja.

```go
person := map[string]string{
    "name":    "Raka",
    "address": "Jakarta",
}

fmt.Println(person["name"])     // akses nilai berdasarkan key
fmt.Println(len(person))        // jumlah key-value
```

**Operasi penting pada Map:**

| Operasi | Contoh |
|---------|--------|
| Akses nilai | `person["name"]` |
| Tambah / ubah | `person["title"] = "Programmer"` |
| Hapus key | `delete(person, "title")` |
| Panjang map | `len(person)` |

---

### 13. If Expression — `if_expression.go`

Percabangan kondisi menggunakan `if`, `else if`, dan `else`.

```go
var age = 15

if age > 18 {
    fmt.Println("Anda sudah dewasa")
} else if age > 12 {
    fmt.Println("Anda masih remaja")
} else {
    fmt.Println("Anda masih anak-anak")
}
```

**If dengan Short Statement** — mendeklarasikan variabel langsung di dalam `if`:

```go
if length := len(name); length > 5 {
    fmt.Println("Nama terlalu panjang")
} else {
    fmt.Println("Nama sudah benar")
}
```

**If dengan Multiple Return Value** — menerima beberapa nilai sekaligus dari fungsi:

```go
func cekStatusGamer(name string) (int, bool, string) {
    return 100, true, "Grandmaster"
}

if score, vip, level := cekStatusGamer("Raka"); vip {
    fmt.Printf("User %s (Skor: %d) sedang online!", level, score)
}
```

> Variabel yang dideklarasikan di short statement hanya berlaku **di dalam blok if** tersebut.

---

### 14. Switch Expression — `switch_expression.go`

Alternatif `if-else` yang lebih rapi untuk kondisi dengan banyak kemungkinan nilai.

```go
var name = "Raka"

// Switch biasa
switch name {
case "Raka":
    fmt.Println("Hello Raka")
case "John":
    fmt.Println("Hello John")
default:
    fmt.Println("Hello World")
}
```

**Switch dengan Short Statement:**

```go
switch length := len(name); length > 5 {
case true:
    fmt.Println("Nama terlalu panjang")
case false:
    fmt.Println("Nama sudah benar")
}
```

**Switch tanpa kondisi** (seperti if-else):

```go
switch {
case name == "Raka":
    fmt.Println("Hello Raka")
case name == "John":
    fmt.Println("Hello John")
default:
    fmt.Println("Hello World")
}
```

---

### 15. For Expression — `for_expression.go`

Go hanya memiliki satu keyword perulangan yaitu `for`, namun bisa digunakan dalam berbagai cara.

**For standar (seperti C/Java):**

```go
for i := 0; i < 10; i++ {
    fmt.Println("Hello World", i)
}
```

**For dengan `range` pada Slice / Array:**

```go
names := []string{"Raka", "John", "Jane"}

for index, value := range names {
    fmt.Println("Index", index, "Value", value)
}
```

**For dengan `range` pada Map:**

```go
newNamesMap := map[string]string{
    "Raka": "Raka",
    "John": "John",
}

for key, value := range newNamesMap {
    fmt.Println("Key", key, "Value", value)
}
```

> Gunakan `_` untuk mengabaikan salah satu dari `index` / `key` jika tidak dibutuhkan, contoh: `for _, value := range names`.

---

### 16. Function — `function.go`

Function adalah blok kode yang bisa dipanggil berulang kali. Di Go, function didefinisikan dengan keyword `func`.

**Function tanpa return value:**

```go
func sayHello(name string) {
    fmt.Println("Hello", name)
}
```

**Function dengan banyak parameter:**

```go
func setIntroduction(name string, age int, address string) {
    fmt.Println("Hello", name, "umur saya", age, "dan saya tinggal di", address)
}
```

**Function dengan named return value:**

```go
func calculate(numberOne int, numberTwo int) (result int) {
    result = numberOne + numberTwo
    return result
}
```

**Memanggil function:**

```go
sayHello("Raka")

age := calculate(10, 20)
setIntroduction("Raka", age, "Jakarta")
```

> **Named return value** membuat kode lebih mudah dibaca karena nama variabel return sudah terdefinisi di signature function.

---

### 17. Variadic Function & Slice Parameter — `function.go`

**Variadic Function** memungkinkan function menerima jumlah argumen yang tidak terbatas menggunakan `...`.

```go
func variadicFunction(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

// Memanggil dengan argumen langsung
variadicFunction(10, 20, 30, 40, 50)

// Memanggil dengan slice menggunakan spread operator (...)
sliceNumbers := []int{10, 20, 30, 40, 50}
variadicFunction(sliceNumbers...)
```

**Slice Parameter** menerima slice secara langsung:

```go
func sliceParameter(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
```

| | Variadic `...int` | Slice `[]int` |
|---|---|---|
| Pemanggilan | `fn(1, 2, 3)` atau `fn(slice...)` | `fn(slice)` |
| Fleksibilitas | Lebih fleksibel | Harus berupa slice |

---

### 18. Function as Value, Parameter & Type Declaration — `function.go`

Di Go, **function adalah first-class citizen** — bisa disimpan ke variabel, dikirim sebagai parameter, dan dideklarasikan sebagai tipe.

**Function as Value** — menyimpan function ke variabel:

```go
func getHello(name string) string {
    return "Hello " + name
}

hello := getHello          // simpan ke variabel
fmt.Println(hello("Raka")) // panggil dari variabel
```

**Function as Parameter** — mengirim function sebagai argumen:

```go
func sayHelloWithFunction(name string, function func(string) string) {
    fmt.Println("Hello", function(name))
}

sayHelloWithFunction("Bagas", getHello)
```

**Function Type Declaration** — membuat alias tipe untuk function:

```go
type Filter func(string) string

func filter(name string, filter Filter) {
    fmt.Println("Hello", filter(name))
}
```

**Anonymous Function** — function tanpa nama, langsung ditulis sebagai argumen:

```go
type Blacklist func(string) bool

isBlacklist("Raka", func(name string) bool {
    return name == "Bagas"
})
```

---

### 19. Closure — `closure.go`

Closure adalah function yang **menangkap (capture) variabel dari scope luar**. Perubahan pada variabel luar akan terpengaruh jika diakses langsung.

```go
counter := 0
increment := func() {
    counter++   // menangkap variabel counter dari luar
    fmt.Println("increment")
}

increment()
increment()
increment()
fmt.Println(counter) // output: 3
```

**Perbedaan capture by reference vs copy:**

```go
counter2 := 0
increment2 := func(c int) { // nilai DISALIN ke parameter
    c++
}

increment2(counter2)
increment2(counter2)
fmt.Println(counter2) // output: 0 (tidak berubah!)
```

| | Capture langsung | Melalui parameter |
|---|---|---|
| Cara kerja | Referensi ke variabel asli | Salinan nilai |
| Efek ke variabel luar | Berubah | Tidak berubah |

---

### 20. Defer — `defer.go`

`defer` digunakan untuk menunda eksekusi sebuah function hingga **function pemanggil selesai**. Sering digunakan untuk cleanup (menutup file, koneksi, dll).

```go
func logging() {
    fmt.Println("Logging")
}

func runApplication() {
    defer logging() // akan dieksekusi SETELAH runApplication() selesai
    fmt.Println("Run Application")
}

// Output:
// Run Application
// Logging
```

> `defer` berjalan seperti **LIFO (Last In, First Out)** jika ada lebih dari satu `defer` dalam satu function.

---

### 21. Panic — `panic.go`

`panic` digunakan untuk **menghentikan eksekusi program** secara paksa saat terjadi kondisi yang tidak bisa ditangani. Program akan berhenti dan menampilkan pesan error.

```go
func runApp(error bool) {
    defer endApp()
    if error {
        panic("Error") // program berhenti di sini
    } else {
        fmt.Println("App Running")
    }
}

func main() {
    runApp(true) // akan memicu panic
}
```

> Meskipun `panic` dipanggil, fungsi yang di-`defer` **tetap akan dieksekusi** sebelum program benar-benar berhenti.

---

### 22. Recover — `recover.go`

`recover` digunakan untuk **menangkap panic** agar program tidak crash. Harus dipanggil di dalam fungsi yang di-`defer`.

```go
func endApp() {
    fmt.Println("End App")
    message := recover() // menangkap pesan dari panic
    fmt.Println("Terjadi Error", message)
}

func runApp(error bool) {
    defer endApp() // recover dipanggil di dalam defer ini
    if error {
        panic("Error")
    } else {
        fmt.Println("App Running")
    }
}

func main() {
    runApp(false) // program berjalan normal
}
```

**Alur kerja Defer + Panic + Recover:**

```
runApp() dipanggil
    → panic terpicu
    → defer endApp() dijalankan
        → recover() menangkap pesan panic
    → program lanjut berjalan (tidak crash)
```
```
PENGGANTI PENGGUNAAN TRY-CATCH
```

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
├── map.go               # Map
├── if_expression.go     # If expression
├── switch_expression.go # Switch expression
├── for_expression.go    # For loop
├── function.go          # Function, variadic, function as value/parameter/type
├── closure.go           # Closure
├── defer.go             # Defer
├── panic.go             # Panic
├── recover.go           # Recover
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
12. `map.go` — Map
13. `if_expression.go` — If expression
14. `switch_expression.go` — Switch expression
15. `for_expression.go` — For loop
16. `function.go` — Function, variadic, function as value/parameter/type, anonymous function
17. `closure.go` — Closure
18. `defer.go` — Defer
19. `panic.go` — Panic
20. `recover.go` — Recover

---

## Referensi

- [Dokumentasi Resmi Go](https://go.dev/doc/)
- [Go Tour (Interaktif)](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
