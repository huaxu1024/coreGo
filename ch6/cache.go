package main

import (
	"sync"
)

var (
	mu sync.Mutex
	mapping = make(map[string]string)
)

func lookUp(key string) string {
	mu.Lock()
	defer mu.Unlock()
	return mapping[key]
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookUpA(key string) string {
	cache.Lock()
	defer cache.Unlock()
	return cache.mapping[key]
}