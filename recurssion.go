package main

import "fmt"

func main() {
	hello(5)
	// fmt.Print(a)
}

func hello(n int) {
	if n <= 1 {
		return
	}
	hello(n - 1)
	fmt.Println(n)
	hello(n - 1)
	fmt.Println(n)
}
