package main

import (
	"flag"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

type Arith int

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func (t *Arith) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}

var addr = flag.String("s", "10.68.214.72:7788", "service address")
var zk = flag.String("zk", "10.73.128.85:2181", "zookeeper URL")
var n = flag.String("n", "10.73.128.85:2181", "Arith")

func main() {
	flag.Parse()

	server := rpcx.NewServer()
	rplugin := &plugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zk},
		BasePath:         "/rpcx",
		Metrics:          metrics.NewRegistry(),
		Services:         make([]string, 1),
		UpdateInterval:   10 * time.Second,
	}
	rplugin.Start()
	server.PluginContainer.Add(rplugin)
	server.RegisterName(*n, new(Arith), "weight=5&state=active")
	server.Serve("tcp", *addr)
}
