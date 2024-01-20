package main

func factorialrecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialrecursive(value-1)
	}
}

func main() {
	println(factorialrecursive(4))

}
