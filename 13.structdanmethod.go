package main

import "fmt"

type pelanggan struct {
	nama, alamat string
	umur         int
}

func (customer pelanggan) hallo(name string) {
	fmt.Println("hello", name, "my name is ", customer.nama)
}

func main() {
	var lana pelanggan
	fmt.Println(lana)

	lana.nama = "lana kurniawan"
	lana.alamat = "indonesia"
	lana.umur = 20

	fmt.Println(lana)

	lana.hallo("ucon")
}
