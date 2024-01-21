package main

import "fmt"

func logging() {
	fmt.Println("selesai ")
}

func run(error bool) {
	defer logging()
	if error {
		panic("error")
	}

}
func main() {
	run(true)

}
