package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"
)

func main() {
	double(10)
	switchD("ssss")
}

func forEach() {
	var remove []func()
	strs := []string{"a", "b", "c", "d"}

	// 最佳实践
	for _, str := range strs {
		remove = append(remove, func() {
			fmt.Println(str)
		})
	}
	for _, re := range remove {
		re()
	}
}

func topSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item]  {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func remove() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	for _, rmdir := range rmdirs {
		rmdir()
	}
}

func tempDirs() []string{
	return []string{}
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return []byte{}, nil
}

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()
	return m[key]
}

func bigSlowOperation() {
	defer tarce("bigSlowOperation")()
	// something
	time.Sleep(10 * time.Second)
}

func tarce(msg string) func () {
	now := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(now))	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result)}()
	return x + x
}

func switchD(str string) {
	switch str {
	case "Psaeds":
	case "Hearts":
	case "Clubs" :
	default:
		panic(fmt.Sprintf("invlid suit %q", str))
	}
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	return nil
}