package main

import (
	"log"
	"os"
)

var v1 int
var v2 string
var v3 map[string]int
var v4 []int

var v5 struct {
}

var v6 func(a int) int

type Base struct {
	id   int
	name string
}

type Product struct {
	//	Base
	pid   int
	pname string
}

type myProduct interface {
	test()
}

func (self Product) test() {
	log.Print("product")
	log.Print(self.pname)
}

func more() (id int, name string) {
	id = 12
	name = "more name"
	return
}

func test() {
	log.Print("go run test")
}

func main() {

	//	go test()
	log.Print("tmp web.go")

	f, err := os.Open("./test")

	defer f.Close()
	if err != nil {
		log.Print("err")
	}
	var myp myProduct
	myp = &Product{1, "ptstring"}
	myp.test()

	id, name := more()
	log.Print(id)
	log.Print(name)

	niming := func() {
		log.Print("niming")
	}
	niming()

	const PI = 3.1514326
	log.Print(PI)

	const zero = 0.0
	log.Print(zero)

	const (
		c1 = iota
		c2
		c3
	)
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
	)
	log.Print(c3)
	log.Print(Friday)
	//	panic(404)
	panic("network error")
}
