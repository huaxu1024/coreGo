package ch8

import (
	"io"
	"log"
	"net"
	"os"
)

func tshy() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCoty(os.Stdout, conn)
}

func mustCoty(dst *os.File, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
