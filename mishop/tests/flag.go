package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/larspensjo/config"
)

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

var Topic = make(map[string]string)
var Other = make(map[string]string)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}

	if cfg.HasSection("topicArr") {
		section, err := cfg.SectionOptions("topicArr")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					Topic[v] = options
				}
			}
		}
	}
	if cfg.HasSection("other") {
		section, err := cfg.SectionOptions("other")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("other", v)
				if err == nil {
					Other[v] = options
				}
			}
		}
	}

	fmt.Println(Topic)
	fmt.Println(Topic["debug"])
	fmt.Println(Other)
}
