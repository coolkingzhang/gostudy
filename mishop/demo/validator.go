package main

// https://github.com/jinuljt/goformvalidator
// https://www.douban.com/note/509189856/
// https://github.com/absoludity/goforms

import "github.com/jinuljt/goformvalidator"
//!!IMPORTANT!!
//因为go-validator的bug，validate tag必须要放在最前面,否则validate是无效的
type Person struct {
    Name string `validate:"min=5,max=10" schema:name"` //姓名长度 5-10个字节
    Age int `validate:"min=1,max=150" schema:age"` //年龄1-150岁
}

person := new(Person)
httpRequest.ParseForm()
if err := goformvalidator.Validate(person, httpRequest.Form); err != nil {
    fmt.Printf("%v validate error due to %s", *person, err)
} else {
    fmt.Printf("%v validate success", *person)
}