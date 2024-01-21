package main

import "fmt"

func logging() {
	fmt.Println("selesai ")
	pesan := recover()
	fmt.Println("terjadi", pesan)
}

func run(error bool) {
	defer logging()
	if error {
		panic("error pada fungsi run")
	}

}
func main() {
	run(true)

}
