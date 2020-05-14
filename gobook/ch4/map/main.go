package main

import "fmt"

func makeTest() {
	ages := make(map[string]int)
	//var ages map[string]int
	fmt.Println(ages == nil)

	age, ok := ages["wewe"]
	if !ok {
		fmt.Println("error")
	} else {
		fmt.Println(age)
	}

	ages["aaa"] = 21
	fmt.Println(len(ages) == 0)
}

func equalMap(x, y map[string]int) bool{
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if a, b := y[k]; !b || a != v {
			return false
		}
	}
	return true
}

var graph = make(map[string]map[string]bool)

func  addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	//fmt.Println(equalMap(map[string]int{"a":0}, map[string]int{"b":242}))
}
