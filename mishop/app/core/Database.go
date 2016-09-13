package core

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine = getEngine()

func getEngine() *xorm.Engine {

	var engine *xorm.Engine
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

	//设置日志信息
	fileLog, err := os.Create("./logs/mysql.log")
	if err != nil {
		println(err.Error())
	}
	engine.Logger().SetLevel(core.LOG_INFO)
	engine.SetLogger(xorm.NewSimpleLogger(fileLog))
	return engine
}
