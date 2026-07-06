package main

import "fmt"
import "math"

func hitungLuas(r float32) float32{
	const Phi = 3.14
	return r * r * math.Phi
}

func kelLingkaran(r float32) float32 {
	const Phi = 3.14
	return 2 * math.Phi * r 
}

func main() {
	fmt.Println(hitungLuas(7))
	fmt.Println(kelLingkaran(7))
}