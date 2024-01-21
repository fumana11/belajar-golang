package main

import "fmt"

type Man struct {
	nama string
}

func (man *Man) Maried() {
	man.nama = "Mr. " + man.nama
}

func main() {
	ucon := Man{"ucon"}

	ucon.Maried()

	fmt.Println(ucon.nama)
}
