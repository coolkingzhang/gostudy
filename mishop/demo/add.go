package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Print(z)
}

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, 1)
	}
	for t := 0; t < 1000000000; t++ {

	}
}
