package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

type User struct {
	Id          int64
	Name        string
	Salt        string
	Age         int
	Passwd      string    `xorm:"varchar(200)"`
	Version     int       `xorm:"version"`
	CreatedTime time.Time `xorm:"created"`
	UpdateTime  time.Time `xorm:"updated"`
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/golang?charset=utf8")
	//xorm reverse mysql root:123456@/golang?charset=utf8 templates/goxorm
	//	engine, err = xorm.NewEngine("sqlite3", "./test.db")

	engine.ShowSQL(true)
	engine.SetMaxIdleConns(1)
	engine.SetMaxOpenConns(16)
	tbMapper := core.NewPrefixMapper(core.SameMapper{}, "go_")
	engine.SetTableMapper(tbMapper)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")

	//启动事务会话
	session := engine.NewSession()
	defer session.Close()

	//设置日志信息
	fileLog, err := os.Create("./mysql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.Logger().SetLevel(core.LOG_INFO)
	engine.SetLogger(xorm.NewSimpleLogger(fileLog))

	//同步数据库结构
	engine.Sync(new(User))
	//	err = engine.Sync(new(User), new(Article))

	//插入用户表
	user := new(User)
	user.Name = "TCL-HENRY"
	user.Salt = "woman"
	user.Passwd = "789"
	user.Age = 118

	//    engine.Table() 指定的临时表名优先级最高
	//    TableName() string 其次
	//    Mapper 自动映射的表名优先级最后

	affected, err := engine.Table("go_user").Insert(user)
	fmt.Print(affected)

	//根据id查询
	user1 := &User{Id: 1}
	has, err := engine.Alias("u").Get(user1)
	fmt.Println(has)
	fmt.Println(user1.Name)

	//列表查询
	userList := make([]User, 0)
	engine.Desc("id").Limit(10, 5).Find(&userList)
	fmt.Print(userList[0].Name)

	//总记录数
	user3 := new(User)
	total, err := engine.Alias("o").Where("o.id >?", 1).And("1=1").Count(user3)
	fmt.Print(total)

	//	//通过SQL
	//	beans := make([]User, 0)
	//	engine.Sql("select * from go_user").Find(&beans)
	//	fmt.Print(beans)

	//	//更新
	//	var user5 User
	//	user5.Passwd = "9898"
	//	user5.Salt = "woman"
	//	engine.Id(1).Update(&user5)

	//更新sql
	//当调用Query时，第一个返回值results为[]map[string][]byte的形式。
	sql := "select * from go_user"
	results, err := engine.Query(sql)
	if err != nil {
		session.Rollback()
		return
	}
	//	fmt.Print(results)
	var pass string
	pass = string(results[0]["passwd"])
	fmt.Print(pass)

	//原生更新sql
	//	sql = "update `userinfo` set username=? where id=?"
	//	res, err := engine.Exec(sql, "xiaolun", 1)

	err = session.Commit()
	if err != nil {
		return
	}

	var userVersion User
	engine.Id(1).Get(&userVersion)
	// SELECT * FROM user WHERE id = ?
	engine.Id(1).Update(&userVersion)
	// UPDATE user SET ..., version = version + 1 WHERE id = ? AND version = ?
}
