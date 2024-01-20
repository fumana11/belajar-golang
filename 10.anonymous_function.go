package main

import "fmt"

type Blacklist func(string) bool

func register(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blacklisted", name)
	} else {
		fmt.Println("wellcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "ucon"
	}

	register("lana", blacklist)
	register("ucon", blacklist)

}
