package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"time"
)

var period = flag.Duration("period", 1 * time.Second, "sleep period")

func sleep() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

const debug = false

func buge() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		// 使用bug
	}
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}