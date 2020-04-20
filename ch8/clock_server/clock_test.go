package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
	"testing"
)

func Test_clock1_conn_max_num(t *testing.T) {
	max := 100
	var wg sync.WaitGroup
	wg.Add(max)
	for i := 0; i < max; i++ {
		go getTcpConnByClock(wg.Done)
	}
	wg.Wait()
}

func getTcpConnByClock(done func()) {
	defer done()
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
}
