package main

import (
	"fmt"
)

func main() {
	var nilai int

	fmt.Print("masukan nilai : ")
	fmt.Scan(&nilai)

	switch nilai {
	case 10:
		fmt.Println("sangat bagus")
	case 8:
		fmt.Println("bagus")
	case 5:
		fmt.Println("cukup")
	default:
		fmt.Println("coba lagi")
	}
}
