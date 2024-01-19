package main

import "fmt"

func main() {
	var nama1 = "lana"
	var nama2 = "lana"

	var hasil bool = nama1 == nama2
	fmt.Println(hasil)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
