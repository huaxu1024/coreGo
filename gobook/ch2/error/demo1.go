package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main () {
	fmt.Println(cwd)
}

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("working directrory = %s", cwd)
}