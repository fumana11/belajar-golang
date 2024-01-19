package main

import "fmt"

func main() {
	// konversi int32 ke int64
	var nilai32 int32 = 12345
	var nilai64 int64 = int64(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// konveri string ke byte
	var nama = "furqon"
	var f byte = nama[0]
	var fstring string = string(f)

	fmt.Println(nama)
	fmt.Println(fstring)
}
