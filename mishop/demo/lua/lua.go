package main

import "github.com/yuin/gopher-lua"
import "fmt"

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}

	//	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}

	lv := L.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println(string(str))
	}
	//	if lv.Type() != lua.LTString {
	//		panic("string required.")
	//	}
}
