package main

import "log"

func main() {
	log.Print("main")
	go run()
}
func run() {
	log.Printf("log")
}
