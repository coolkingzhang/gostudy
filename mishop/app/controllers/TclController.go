package controllers

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Bird struct {
	Name           string
	LiftExpectance int
}

func (b *Bird) Flay() {
	fmt.Print("I am flying ...... ")
}

type TclController struct {
}

func (self TclController) Test(c *gin.Context) {
	defer fmt.Print(" defer")
	fmt.Print(bibao(123))
	one, two, three := morer()
	fmt.Print(one + two + three)
	fans()

	const PI string = "pid"
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Print(Saturday)
	fmt.Print(PI)

}

//闭包
func bibao(par int) int {
	f := func(x int) int {
		return x
	}
	return f(par)
}

//多返回值
func morer() (one string, two string, three string) {
	one = "one"
	two = "two"
	three = "three"
	return
}

//反射条件函数
func fans() {
	//反射函数
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
