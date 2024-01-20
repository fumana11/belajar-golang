package main

import "fmt"

func sumall(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	fmt.Println(sumall(10, 10, 10))
	fmt.Println(sumall(15, 15, 15))

	numbers := []int{10, 15, 15}
	fmt.Println(sumall(numbers...))

}
