package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1 * time.Second)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {

}