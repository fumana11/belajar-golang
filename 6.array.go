package main

import "fmt"

func main() {
	var buah [4]string
	buah[0] = "apel"
	buah[1] = "anggur"
	buah[2] = "jambu"
	buah[3] = "manggis"

	for i := 0; i < len(buah); i++ {
		fmt.Println("buah [", i, "] :", buah[i])
	}
}
