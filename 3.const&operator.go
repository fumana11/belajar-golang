package main

import "fmt"

func main() {
	// constan
	const tinggi = 100
	fmt.Println("nilai tinggi = ", tinggi)

	//operasi bilangan
	var jumlah = ((2 + 2) * 2) / 4
	var uji = jumlah != 3
	fmt.Printf("jumlah = %d (%t) \n", jumlah, uji)

	//logika matematika
	var kiri = true
	var kanan = false

	kirinkanan := kiri && kanan
	kiriokanan := kiri || kanan
	notkiri := !kiri

	fmt.Printf("kiri dan kanan adalah %t \n", kirinkanan)
	fmt.Printf("kiri atau kanan adalah %t \n", kiriokanan)
	fmt.Printf("bukan kiri adalah %t \n", notkiri)
}
