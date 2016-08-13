package core

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine = getEngine()

func getEngine() *xorm.Engine {
	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/golang?charset=utf8")
	if err != nil {
		fmt.Print(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(16)
	engine.SetMaxIdleConns(2)
	return engine
}
