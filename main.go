package main

import "fmt"

func hitungLuas(r int) float32{
	var phi float32
	if r%7 == 0 {
		phi = 22.0 / 7.0
	} else{
		phi = 3.14
	}
	return float32(r) * float32(r) * phi
}

func kelLingkaran(r int) float32 {
	var phi float32
	if r%7 == 0 {
		phi = 22.0 / 7.0
	} else{
		phi = 3.14
	}
	return 2 * phi * float32(r) 
}

func main() {
	var r int
	fmt.Print("Masukan jari-jari lingkaran : ")
	fmt.Scan(&r)

	keliling := kelLingkaran(r)
	luas := hitungLuas(r)

	fmt.Println("======= Hasil =======")
	fmt.Println("Jari - Jari : ", r)
	fmt.Println("Luas Lingkarang : ", luas)
	fmt.Println("Keliling Lingkaran : ", keliling)
}