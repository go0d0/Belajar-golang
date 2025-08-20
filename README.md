#  Perjalanan Belajar Go (Golang)

Selamat datang di repositori perjalanan belajar Go saya! Repositori ini adalah kumpulan kode yang saya tulis sambil mempelajari berbagai konsep dalam bahasa pemrograman Go, mulai dari dasar hingga konkurrensi. Semua contoh kode ada di dalam file `main.go`.

## 📚 Daftar Isi

1.  [Dasar-dasar Go](#1-dasar-dasar-go)
    -   [Variabel](#variabel)
    -   [Kondisi `if-else`](#kondisi-if-else)
    -   [Kondisi `switch`](#kondisi-switch)
    -   [Looping (`for`, `break`, `continue`)](#looping-for-break-continue)
2.  [Fungsi](#2-fungsi)
    -   [Fungsi Dasar](#fungsi-dasar)
    -   [Fungsi sebagai Parameter (Callback)](#fungsi-sebagai-parameter-callback)
    -   [Fungsi Anonim](#fungsi-anonim)
    -   [Closure](#closure)
3.  [Struktur Data](#3-struktur-data)
    -   [Array](#array)
    -   [Slice](#slice)
    -   [Map](#map)
4.  [Struct](#4-struct)
    -   [Struct Dasar](#struct-dasar)
    -   [Embedded Struct](#embedded-struct)
    -   [Method pada Struct](#method-pada-struct)
5.  [Pointer](#5-pointer)
6.  [Interface](#6-interface)
    -   [Deklarasi dan Implementasi Interface](#deklarasi-dan-implementasi-interface)
    -   [Interface Kosong (`interface{}`)](#interface-kosong-interface)
    -   [Type Assertion](#type-assertion)
7.  [Penanganan Error](#7-penanganan-error)
    -   [Error Dasar (`errors.New`)](#error-dasar-errorsnew)
    -   [Error Formatting (`fmt.Errorf` dan `%w`)](#error-formatting-fmterrorf-dan-w)
    -   [Pengecekan Error (`errors.Is`)](#pengecekan-error-errorsis)
    -   [Pengecekan Tipe Error Spesifik (`errors.As`)](#pengecekan-tipe-error-spesifik-errorsas)
    -   [Custom Error](#custom-error)
8.  [Concurrency](#8-concurrency)
    -   [Goroutine dan `sync.WaitGroup`](#goroutine-dan-syncwaitgroup)
    -   [Channel](#channel)
    -   [Select](#select)
9.  [Package dan Module](#9-package-dan-module)

---

### 1. Dasar-dasar Go

#### Variabel
Deklarasi dan inisialisasi variabel.
```go
umur := 12      // Deklarasi singkat
orang = "Person" // Variabel package-level
```

#### Kondisi `if-else`
Pengecekan kondisi untuk mengontrol alur program.
```go
if orang == "Orang" {
    fmt.Println("halo", orang)
} else if orang == "Person" {
    fmt.Println("halo", orang)
} else {
    fmt.Println("tidak dikenali")
}
```

#### Kondisi `switch`
Alternatif dari `if-else` untuk pengecekan banyak kondisi pada satu variabel.
```go
switch umur {
case 20:
    fmt.Println("umur", umur)
case 12:
    fmt.Println("umur", umur)
}
```

#### Looping (`for`, `break`, `continue`)
Go hanya memiliki satu jenis loop, yaitu `for`.
```go
// Loop dasar
for x := 1; x < 6; x++ {
    fmt.Println("loop ke:", x)
}

// Loop dengan break
for y := 2; y < 4; y++ {
    if y == 3 {
        break
    }
    fmt.Println("tes break:", y)
}

// Loop dengan continue
for a := 0; a < 5; a++ {
    if a == 1 {
        continue
    }
    fmt.Println("nilai continue:", a)
}
```

### 2. Fungsi

#### Fungsi Dasar
Fungsi yang dipanggil dari `main` dengan parameter.
```go
func tes(nama string, jalan func(string) string) string {
    if nama == "Abc" {
        return "halo " + nama
    }
    return jalan("data tidak sesuai")
}
```

#### Fungsi sebagai Parameter (Callback)
Fungsi yang menjadi argumen untuk fungsi lain.
```go
func tes1(data string) string { // callback function
    return data + ", mohon diperiksa lagi"
}

// Memanggil fungsi `tes` dengan `tes1` sebagai callback
fmt.Println(tes("John doe", tes1))
```

#### Fungsi Anonim
Fungsi yang dideklarasikan tanpa nama, seringkali langsung dieksekusi.
```go
// Dideklarasikan dan langsung dipanggil
angka := func(a, b int) int {
    return a * b
}(5, 4)
fmt.Println("nilai anonymous angka:", angka)

// Disimpan dalam variabel
karakter := func() string {
    return "hallo string anonymous"
}
fmt.Println(karakter())
```

#### Closure
Fungsi yang "mengingat" lingkungan tempat ia dibuat. Closure dapat mengakses variabel di luar body-nya.
```go
func bClosure() func() int {
	nilai := 0
	return func() int { //closure
		nilai++
		return nilai
	}
}

// Setiap panggilan akan mengincrement `nilai` yang sama
penghitung := bClosure()
fmt.Println(penghitung()) // Output: 1
fmt.Println(penghitung()) // Output: 2
```

### 3. Struktur Data

#### Array
Kumpulan elemen dengan tipe data yang sama dan panjang yang tetap.
```go
var a = [2]int{23, 89}
fmt.Println("belajar array: ", a)
```

#### Slice
Versi dinamis dan lebih fleksibel dari array.
```go
mm := []int{11, 1234, 3}
fmt.Println("belajar slice: ", mm)

// Menambah elemen baru
mm = append(mm, 999)

// Panjang (jumlah elemen) dan Kapasitas (ukuran alokasi memori)
fmt.Println("digunakan (len): ", len(mm))
fmt.Println("tersedia (cap): ", cap(mm))

// Looping slice dengan range
for i, v := range mm {
    fmt.Println("index: ", i, "valuenya :", v)
}
```

#### Map
Kumpulan pasangan key-value yang tidak berurutan.
```go
var ahay map[string]int = map[string]int{
    "satu": 1,
    "dua":  2,
}

ahay["tiga"] = 3     // Menambah data baru
delete(ahay, "satu") // Menghapus data
fmt.Println(ahay)
```

### 4. Struct

#### Struct Dasar
Tipe data kustom yang berisi kumpulan *field*.
```go
type tesLagi struct {
    nama   string
    jumlah int
}
lg := tesLagi{"budi", 3}
fmt.Println(lg)
```

#### Embedded Struct
Struct yang disematkan di dalam struct lain.
```go
type coba struct {
    nama   string
    jumlah int
}

type tinggi struct {
    coba      // embedded struct
    kebenaran bool
}

var opini tinggi = tinggi{
    coba: coba{"Qwerty", 5},
    kebenaran: true,
}
```

#### Method pada Struct
Fungsi yang terikat pada tipe data struct tertentu.
```go
type oring struct {
	hidup    bool
	bernafas int
}

func (x *oring) cobaAh() {
	x.bernafas++
}

// Pemanggilan method
var mq oring = oring{true, 5}
mq.cobaAh() // bernafas menjadi 6
```

### 5. Pointer
Pointer adalah variabel yang menyimpan alamat memori dari variabel lain.
```go
qkl := 10
var zzz *int = &qkl // zzz menunjuk ke alamat memori qkl
fmt.Println(*zzz)   // Mengakses nilai qkl melalui pointer (dereferencing)
```

### 6. Interface

#### Deklarasi dan Implementasi Interface
Interface adalah sebuah tipe data abstrak yang mendefinisikan sekumpulan *method*. Sebuah tipe data dianggap mengimplementasikan interface jika memiliki semua method yang didefinisikan oleh interface tersebut.
```go
// Deklarasi interface
type bLajar interface {
	jalankan(data string, data1 string)
}

// Struct yang akan mengimplementasikan interface
type manggil struct {
	nama string
	umur int
}

// Implementasi method, secara otomatis `manggil` memenuhi kontrak `bLajar`
func (x *manggil) jalankan(data string, data1 string) {
	fmt.Println(x.nama, x.umur)
	fmt.Println("ini data1 interface: ", data, "ini adalah data2 interface: ", data1)
}
```

#### Interface Kosong (`interface{}`)
Tipe data yang dapat menampung nilai dari tipe data apa pun.
```go
var kosong interface{}
kosong = 5 // Bisa diisi integer, string, bool, dll.

// Pengecekan tipe data dengan switch type
switch pp := kosong.(type) {
case string:
    fmt.Println("tipe adalah string: ", pp)
case int:
    fmt.Println("tipe adalah integer: ", pp)
}
```

#### Type Assertion
Mekanisme untuk mendapatkan nilai konkret dari sebuah variabel interface.
```go
var data interface{} = "hello"
xx := data.(string) // Assert bahwa `data` adalah string
fmt.Println(xx, "dunia!")
```

### 7. Penanganan Error

#### Error Dasar (`errors.New`)
Cara paling sederhana untuk membuat error.
```go
func salah(data1, data2 float64) (float64, error) {
	if data2 == 0 {
		return 0, errors.New("data tidak valid untuk operasi")
	}
	return data1 / data2, nil
}
```

#### Error Formatting (`fmt.Errorf` dan `%w`)
Membuat error baru sambil "membungkus" (wrapping) error yang sudah ada. Ini berguna untuk menambah konteks pada error tanpa menghilangkan error aslinya.
```go
func adaYgSalah() error {
	waduh := errors.New("tes salah")
	return fmt.Errorf("waduh ada yang salah, salah adalah: %w", waduh)
}
```

#### Pengecekan Error (`errors.Is`)
Memeriksa apakah sebuah error dalam rantai error (error chain) sama dengan error target.
```go
var ngawur = errors.New("kesalahan fatal")

func fang() error {
	return fmt.Errorf("terdapat error %w", ngawur)
}

// Pengecekan di main
// halo := fang()
// if errors.Is(halo, ngawur) {
//     fmt.Println("berhasil cek dengan errors.is()")
// }
```

#### Pengecekan Tipe Error Spesifik (`errors.As`)
Memeriksa apakah sebuah error dalam rantai error cocok dengan tipe tertentu, dan jika cocok, nilainya akan di-assign ke variabel.
```go
type kesalahan struct {
	code int
}

func (k kesalahan) Error() string {
	return "hanya tes"
}

func memanggil() error {
	ambil := kesalahan{12}
	return fmt.Errorf("kesalahan program, dengan kode: %w", ambil)
}

// Pengecekan di main
// erur := memanggil()
// var kondisi kesalahan
// if errors.As(erur, &kondisi) {
//     fmt.Println("kesalahan dengan kode", kondisi.code)
// }
```

#### Custom Error
Membuat tipe error sendiri dengan struct yang mengimplementasikan interface `error`.
```go
type custom struct {
	bagian string
	pesan  string
}

func (c *custom) Error() string {
	return fmt.Sprintf("fatal error terjadi di %s, dengan pesan error %s", c.bagian, c.pesan)
}

func menjalankan(data string) error {
	if len(data) < 6 {
		return &custom{"username", "kurang dari 6"}
	}
	return nil
}
```

### 8. Concurrency

#### Goroutine dan `sync.WaitGroup`
Goroutine adalah sebuah *lightweight thread* yang dikelola oleh Go runtime. `sync.WaitGroup` digunakan untuk menunggu semua goroutine selesai dieksekusi.
```go
func berjalanlah(data string, wi *sync.WaitGroup) {
	defer wi.Done() // Memberi tahu WaitGroup bahwa goroutine ini selesai
	// ...
}

func main() {
    var ww sync.WaitGroup
    ww.Add(1) // Menambah satu task ke counter WaitGroup
    go berjalanlah("halo dari goroutine", &ww)
    ww.Wait() // Menunggu sampai counter menjadi nol
    fmt.Println("output setelah goroutine")
}
```

#### Channel
Channel digunakan oleh goroutine untuk berkomunikasi dan sinkronisasi.
```go
// Unbuffered Channel (sinkron)
ch := make(chan string)
go func() {
    ch <- "tes123" // Mengirim data ke channel
}()
fmt.Println(<-ch) // Menerima data dari channel

// Buffered Channel (asinkron)
cihuy := make(chan int, 2) // Kapasitas 2
cihuy <- 5
cihuy <- 3
fmt.Println(<-cihuy)
fmt.Println(<-cihuy)
```

#### Select
`select` memungkinkan sebuah goroutine untuk menunggu beberapa operasi komunikasi (channel). `select` akan memblokir sampai salah satu case-nya bisa berjalan.
```go
ceha := make(chan string)
ceha1 := make(chan string)

// ... (dua goroutine mengirim data ke ceha dan ceha1)

select {
case msg := <-ceha:
    fmt.Println("pesan ch 1 : ", msg)
case msg1 := <-ceha1:
    fmt.Println("pesan ch 2 : ", msg1)
}
```

### 9. Package dan Module
Memanggil fungsi dari package lain yang ada di dalam modul.
```go
import "coba/tes2"

// ... di dalam main
tes2.Halo() // Memanggil fungsi Halo() dari package tes2
```
