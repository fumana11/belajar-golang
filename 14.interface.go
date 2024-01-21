package main

import "fmt"

type Person struct {
	nama string
}

type HasName interface {
	getname() string
}

func (person Person) getname() string {
	return person.nama
}

func sayhello(value HasName) {
	fmt.Println("hello", value.getname())
}
func main() {
	orang := Person{nama: "ucon"}
	sayhello(orang)

}
