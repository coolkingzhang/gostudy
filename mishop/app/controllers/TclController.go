package controllers

import (
	"fmt"
	"mishop/app/models"
	"mishop/app/utils"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinuljt/goformvalidator"
)

type Person struct {
	Name string `validate:"min=5,max=10" schema:name"` //姓名长度 5-10个字节
	Age  int    `validate:"min=1,max=150" schema:age"` //年龄1-150岁
}

type Bird struct {
	Name           string
	LiftExpectance int
}

func (b *Bird) Flay() {
	fmt.Print("I am flying ...... ")
}

type TclController struct {
}

func (self TclController) From(c *gin.Context) {
	person := new(Person)
	c.Request.ParseForm()
	if err := goformvalidator.Validate(person, c.Request.Form); err != nil {
		fmt.Printf("%v validate error due to %s", *person, err)
	} else {
		fmt.Printf("%v validate success", *person)
	}
}

func (self TclController) Test(c *gin.Context) {
	//	defer fmt.Print(" defer")
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

	user := models.UserModel{}.GetUserAll()
	c.JSON(http.StatusUnauthorized, user)

	str := utils.EncryptUtil{}.Md5("123456")
	fmt.Print(str)

	str1 := utils.EncryptUtil{}.Hmac("123456", "key")
	fmt.Print(str1)
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
