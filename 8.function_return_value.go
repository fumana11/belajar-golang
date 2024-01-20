package main

import "fmt"

func gethello(nama string) string {
	hello := "hello " + nama
	return hello
}

func main() {
	hasil := gethello("lana")
	fmt.Println(hasil)
	fmt.Println(gethello("ucon"))
}
