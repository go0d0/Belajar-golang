package main

import "fmt"

var orang string

func tes(nama string) string { //fungsi diluar main
	if nama == "jokowi" {
		return "halo " + nama
	} //multiple return
	return "anda siapa?"
}

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

	fmt.Println(tes("jokowis")) //manggil fungsi tes
}
