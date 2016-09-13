package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

// https://github.com/syndtr/goleveldb/blob/master/README.md

func main() {
	db, err := leveldb.OpenFile("./leveldb", nil)
	defer db.Close()

	err = db.Put([]byte("tclkey"), []byte("welcome to tcl"), nil)
	if err != nil {
		fmt.Print(err)
	}

	data, err := db.Get([]byte("tclkey"), nil)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(data))

	err = db.Delete([]byte("tclkey"), nil)
	if err != nil {
		fmt.Print(err)
	}
}
