package service

import "fmt"

func Test() {
	fmt.Print("test")
}

func TestCase() {
	fmt.Print("case")
}

type Product interface {
	Echo() string
}
