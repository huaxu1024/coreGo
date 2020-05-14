package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err == nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	return []string{}, nil
}

func error01() (string, error){
	_, err := http.Get("url")
	if err != nil {
		return "error", err
	}
	return "", nil
}

func wait(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	os.RemoveAll("")
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func demo1() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func squares() func() int {
	var x int
	return func() int {
		x ++
		return x * x
	}
}

func maixxn() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	x := squares()
	fmt.Println(x())
}