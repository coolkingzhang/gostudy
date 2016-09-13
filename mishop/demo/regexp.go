package main

import (
	"fmt"
	"regexp"
)

var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)
var digitsRegexp = regexp.MustCompile(`(\d+)\D+(\d+)`)

func main() {
	someString := "1000abcd123"
	fmt.Println(digitsRegexp.FindStringSubmatch(someString))
	fmt.Printf("%+v", myExp.FindStringSubmatch("1234.5678.9"))
}
