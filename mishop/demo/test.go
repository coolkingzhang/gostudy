package main

import (
	"fmt"
	"mishop/app/core"
)

//import引用 的是 文件家的路劲，但使用的时候是使用 包名的路径

func main() {
	body := core.HttpGet("http://www.tcl.com/")
	fmt.Print(body)
}
