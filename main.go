package main

import "fmt"

var orang string

func berjalan(f func(string)string)string {


return f("tt")
}


func tes1(data string)string{  // callback function
return data+ "ngawur";
}


func tes(nama string, jalan func(string)string) string { //fungsi diluar main
	if nama == "jokowi" {
		return "halo " + nama
	} //multiple return
	return jalan("data tidak sesuai");  //callback
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

	fmt.Println(tes("jokowis", tes1)) //manggil fungsi tes



	angka := func (a,b int) int{
return a*b
	}(5,4)
	fmt.Println("nilai anonymous angka:",angka);


	karakter := func()string{

		return "hallo string anonymous";
}
fmt.Println(karakter());

simpan := berjalan(func (data string) string{

return "halo "+data
});

fmt.Println(simpan);

}
