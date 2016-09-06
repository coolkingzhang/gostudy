package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type User struct {
	Id          int64
	Name        string
	Salt        string
	Age         int
	Passwd      string    `xorm:"varchar(200)"`
	CreatedTime time.Time `xorm:"created"`
	UpdateTime  time.Time `xorm:"updated"`
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/golang?charset=utf8")
	engine.ShowSQL(true)

	f, err := os.Create("./sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetMaxIdleConns(1)
	engine.SetMaxOpenConns(8)
	engine.SetLogger(xorm.NewSimpleLogger(f))
	tbMapper := core.NewPrefixMapper(core.SameMapper{}, "go_")
	engine.SetTableMapper(tbMapper)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	engine.Sync(new(User))

	user := new(User)
	user.Name = "zhangzhihe"
	user.Salt = "man"
	user.Passwd = "123456789"
	user.Age = 15
	affected, err := engine.Insert(user)
	fmt.Print(affected)

	//	user1 := &User{Id: 1}
	//	has, err := engine.Get(user1)
	//	fmt.Println(has)
	//	fmt.Println(user1.Name)

}
