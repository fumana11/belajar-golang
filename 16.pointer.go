package main

import "fmt"

type alamat struct {
	kota, provinsi, negara string
}

func main() {
	alamat1 := alamat{"banyuwangi", "jawa timur", "indonesia"}
	alamat2 := &alamat1

	alamat2.kota = "bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
