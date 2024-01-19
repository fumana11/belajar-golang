package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)
}
