package main

import (
	"fmt"
)

func main() {
	var nilai int

	fmt.Print("masukan nilai : ")
	fmt.Scan(&nilai)

	if nilai == 10 {
		fmt.Println("sangat bagus")
	} else if nilai < 10 && nilai >= 5 {
		fmt.Println("bagus")
	} else if nilai < 5 && nilai > 1 {
		fmt.Println("cukup")
	} else {
		fmt.Println("coba lagi")
	}
}
