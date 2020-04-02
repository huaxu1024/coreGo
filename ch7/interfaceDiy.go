package main

import (
	"io"
)

type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type ReadWriterCloser interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	io.Closer
}

func impl() {
}
