package main

import (
	"log"

	"github.com/c4pt0r/ini"
)

var conf = ini.NewConf("config.ini")

var (
	v1 = conf.String("topicArr", "addr", "v1")
	//	v2 = conf.Int("topicArr", "login", 0)
	v2 = conf.String("topicArr", "login", "my login")
)

func main() {
	conf.Parse()

	log.Println(*v1, *v2)
}
