package main

import (
	"errors"
	"fmt"
)

// error with .as()
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

// error is()
var ngawur = errors.New("salah sekali")

func fang() error {
	return fmt.Errorf("terdapat error %w", ngawur)
}

// fmt ErrorF
func adaYgSalah() error {
	waduh := errors.New("tes salah")
	return fmt.Errorf("waduh ada yang salah, salah adalah: %w", waduh)

}

// fungsi error
func salah(data1, data2 float64) (float64, error) {
	if data2 == 0 {
		return 0, errors.New("data tidak valid untuk operasi ")
	}

	return data1 / data2, nil
}

// manggil interface

type manggil struct {
	nama string
	umur int
}

func (x *manggil) jalankan(data string, data1 string) {
	fmt.Println(x.nama, x.umur)
	fmt.Println("ini data1 interface: ", data, "ini adalah data2 interface: ", data1)
}

// learn make interface
type bLajar interface {
	jalankan(data string, data1 string)
}

type oring struct {
	hidup    bool
	bernafas int
}

func (x *oring) cobaAh() {
	x.bernafas++
}

func bClosure() func() int {

	nilai := 0

	return func() int { //closure

		nilai++
		return nilai // nested return
	}

}

// tes logika pemanggilan ()() atau pemanggilan func return untuk dapat hasil
func keDua() func() string {

	return func() string {
		return "halo nested" //multiple return
	}
}

var orang string

func berjalan(f func(string) string) string {

	return f("tt")
}

func tes1(data string) string { // callback function
	return data + "ngawur"
}

func tes(nama string, jalan func(string) string) string { //fungsi diluar main
	if nama == "jokowi" {
		return "halo " + nama
	} //multiple return
	return jalan("data tidak sesuai") //callback
}

// kode utama  @@@@@
func main() {

	umur := 12 // variabe umur

	orang = "tono" // variabel orang

	//cek value orang == "jaki"  atau == "tono" atau yang lain
	if orang == "jaki" {
		fmt.Println("halo", orang)
	} else if orang == "tono" {

		fmt.Println("halo", orang)

	} else {

		fmt.Println("ngawur")
	}

	// pengecekan kondisi umur
	switch umur {
	case 20:
		fmt.Println("umur", umur)
	case 12:
		fmt.Println("umur", umur)
	}

	for x := 1; x < 6; x++ { //looping sampai 5

		fmt.Println("loop ke:", x)
	}

	for y := 2; y < 4; y++ { //tes break
		if y == 3 {
			break

		}
		fmt.Println("tes break:", y)
	}

	for a := 0; a < 5; a++ { // tes continue
		if a == 1 {
			continue
		}
		fmt.Println("nilai continue:", a)
	}

	fmt.Println(tes("jokowis", tes1)) //manggil fungsi tes

	angka := func(a, b int) int {
		return a * b
	}(5, 4)
	fmt.Println("nilai anonymous angka:", angka)

	karakter := func() string {

		return "hallo string anonymous"
	}
	fmt.Println(karakter())

	simpan := berjalan(func(data string) string {

		return "halo " + data
	})

	fmt.Println(simpan)

	fmt.Println(bClosure()())
	fmt.Println(keDua()())

	// array
	var a = [2]int{23, 89}
	fmt.Println("belajar array: ", a) // print array

	// slice

	mm := []int{11, 1234, 3}           // slice
	fmt.Println("belajar slice: ", mm) // print slice

	fmt.Println(append(mm, 999)) // append

	fmt.Println("digunakan: ", len(mm)) // len
	fmt.Println("\n")                   // new line
	qw := append(mm, 66)                // apend
	fmt.Println("tersedia: ", cap(qw))  // cap

	// for loop using range
	for i, v := range mm {
		fmt.Println("index: ", i, "valuenya :", v)
	}

	// map in go
	var ahay map[string]int = map[string]int{
		"satu": 1,
		"dua":  2,
	}

	ahay["tiga"] = 3     // add key value "tiga" : 3
	delete(ahay, "satu") // hapus key "satu"
	val := ahay          // simpan ke val

	fmt.Println(val) // print semua

	// learn how to make struct in golang
	type tesLagi struct {
		nama   string
		jumlah int
	}
	lg := tesLagi{"budi", 3}

	fmt.Println(lg)

	// learn how to make embedded struct
	type coba struct {
		nama   string
		jumlah int
	}

	type tinggi struct {
		coba
		kebenaran bool
	}

	var opini tinggi = tinggi{

		coba: coba{
			"joka", 5,
		}, kebenaran: true,
	}

	fmt.Println(opini)

	// method struct

	var mq oring = oring{true, 5}
	fmt.Println(mq)
	fmt.Println("sebelum increment bernafas: ", mq.bernafas)
	mq.cobaAh()
	fmt.Println("setelah increment bernafas: ", mq.bernafas)

	// basic pointer

	qkl := 10

	var zzz *int = &qkl

	if zzz == &qkl {
		fmt.Println(*zzz)
	}

	//run interface
	yyy := manggil{
		"budi", 10,
	}

	yyy.jalankan("halo", "disana")

	// interface kosong

	var kosong interface{}

	kosong = 5 // insert ke interface

	switch pp := kosong.(type) { // cek tipe data
	//cek kondisi apakah tipe data sesuai dengan tipe data variabel pp ?
	case string:
		fmt.Println("tipe adalah string: ", pp) // tampilin value
	case bool:
		fmt.Println("tipe adalah boolean: ", pp) // tampilin value
	case int:
		fmt.Println("tipe adalah integer: ", pp) // tampilin value
	}

	// tes manggil error
	// kalau parameter kedua diganti jadi 3,
	// maka variabel hasil akan berisi,
	// 3.3333333333333335
	// hasil, err := salah(10, 0)

	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(hasil)

	// manggil fmt errorf
	fmt.Println(adaYgSalah())

	//cek dengan errors.As()

	//  erur := memanggil()

	//var kondisi kesalahan

	//if errors.As(erur, &kondisi) {
	//	fmt.Println("kesalahan dengan kode", kondisi.code)
	//	return
	//}

	// cek dengan errors.is()
	//halo := fang()

	// if errors.Is(halo, ngawur) {
	//	fmt.Println("berhasil cek dengan errors.is()")
	//	return
	//}

}
