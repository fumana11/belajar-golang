package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil aplikasi")
}

func run() {
	defer logging()
	fmt.Println("run aplikasi")
}
func main() {
	run()

}
