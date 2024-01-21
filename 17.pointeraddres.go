package main

import "fmt"

type alamat struct {
	kota, provinsi, negara string
}

func merubahalamat(Alamat *alamat) {
	Alamat.kota = "Indonesia"
}

func main() {
	addresses := alamat{}
	merubahalamat(&addresses)
	fmt.Println(addresses)
}
