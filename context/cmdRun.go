package main

import (
	"flag"
	"fmt"
)

var (
	serv = flag.String("service", "hello_service", "service name")
	host = flag.String("host", "localhost", "listening host")
	port = flag.String("port", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

func main() {
	flag.Parse()
	fmt.Printf("serv: %s, host: %s, port: %s, reg: %s", *serv, *host, *port, *reg)
}