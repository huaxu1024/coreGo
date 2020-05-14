package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Move struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"'`
	Actors []string
}

var movies = []Move{
	{Title: "Casblanca", Year: 1942, Color: true,
		Actors: []string{"a", "b", "c"}},
	{Title: "asdawew", Year: 194342, Color: false,
		Actors: []string{"a1", "b2", "c2"}},
	{Title: "gggggg", Year: 194442, Color: true,
		Actors: []string{}},
}

var title []struct{ Title string}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(title)
}
