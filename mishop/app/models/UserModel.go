package models

import (
	//	"fmt"
	"mishop/app/core"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

type UserModel struct {
}

func (self UserModel) Insert(user User) int64 {
	affected, _ := core.Engine.Insert(user)
	return affected
}

func (self UserModel) GetUser(id int) User {
	var user User
	core.Engine.Id(id).Get(&user)
	return user
}
