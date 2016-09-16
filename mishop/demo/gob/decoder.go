package main

import (
	"encoding/gob"
	"fmt"
	"os"
)
type User struct {
	Id   int
	Name string
}

func (this *User) Say() string {
	return this.Name + ` hello world ! `
}

var u []User
file, err := os.Open("./mygo/gob")
if err != nil {
    fmt.Println(err)
}
dec := gob.NewDecoder(file)
err2 := dec.Decode(&u)

if err2 != nil {
    fmt.Println(err2)
    return
}

for _ , user := range u {
    fmt.Println(user.Id)
    fmt.Println(user.Say())
}