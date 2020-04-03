package main

import (
	"encoding/json"
	"fmt"
	"log"
)


func marshal() {

	skill := make(map[string]float64)
	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	level := make(map[string]interface{})
	level["web"] = "Good"
	level["server"] = 90
	level["tool"] = nil

	user := User{
		Name: "huaxu",
		Age: 25,
		Roles: []string{"Owner", "Master"},
		Skill: skill,
		MAccount: Account{
			Email:    "tshy0425@hotmail.com",
			Nickname: "rookie",
			Password: "123456",
			Money:    12.45,
			Extra: 	  []interface{}{1,"a", 23.232, "b"},
		},
		Level: level,
	}

	marshal, err := json.Marshal(user)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(marshal))
}
