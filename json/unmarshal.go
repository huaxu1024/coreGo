package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

var jsonStringAccount string = `{
	"email":"tshy0425@hotmail.com",
	"password":"123",
	"nickname":"rookie",
	"money":"12.45",
	"Extra":[1,"a",23.232,"b"]
}`

func Decode(r io.Reader)(a *Account, err error)  {
	a = new(Account)
	json.NewDecoder(r).Decode(a)
	return
}

func main() {
	account := Account{}
	err := json.Unmarshal([]byte(jsonStringAccount), &account)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("Unmarshal: %+v\n", account)

	user, err := Decode(strings.NewReader(jsonStringAccount))
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("decode: %#v\n",user)
}


func DecodeLoginUser(r io.Reader) (u *LoginUser, err error) {
	u = new(LoginUser)
	if err = json.NewDecoder(r).Decode(u); err != nil{
		return
	}
	switch t := u.UserName.(type) {
	case string:
		u.UserName = t
	case float64:
		u.UserName = int64(t)
	}
	return
}

var jsonString string = `{
        "things": [
            {
                "name": "Alice",
                "age": 37
            },
            {
                "city": "Ipoh",
                "country": "Malaysia"
            },
            {
                "name": "Bob",
                "age": 36
            },
            {
                "city": "Northampton",
                "country": "England"
            }
        ]
    }`

func decode(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]map[string]interface{}
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range data["things"] {
		item := data["things"][i]
		if item["name"] != nil {
			persons = addPerson(persons, item)
		} else {
			places = addPlace(places, item)
		}

	}
	return
}

func addPerson(persons []Person, item map[string]interface{}) []Person {
	name := item["name"].(string)
	age := item["age"].(float64)
	person := Person{name, int(age)}
	persons = append(persons, person)
	return persons
}

func addPlace(places []Place, item map[string]interface{}) []Place {
	city := item["city"].(string)
	country := item["country"].(string)
	place := Place{City:city, Country:country}
	places = append(places, place)
	return places
}

// 混合结构
func decodePersonPlace(jsonStr []byte) (persons []Person, places []Place) {
	var data map[string][]Mixed
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", data["things"])
	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{Name: item.Name, Age: item.Age})
		} else {
			places = append(places, Place{City: item.city, Country: item.Country})
		}
	}
	return
}

// json.RawMessage
func decodeRawMessage(jsonStr []byte)(persons []Person, places []Place){
	var data map[string][]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil{
		fmt.Println(err)
		return
	}

	for _, item := range data["things"]{
		persons = addPersonRaw(item, persons)
		places = addPlaceRaw(item, places)
	}
	return
}

func addPersonRaw(item json.RawMessage, persons []Person) []Person {
	person := Person{}
	if err := json.Unmarshal(item, &person); err != nil {
		fmt.Println(err)
	} else {
		if person != *new(Person){
			persons = append(persons, person)
		}
	}
	return persons
}

func addPlaceRaw(item json.RawMessage, places []Place) []Place {
	place := Place{}
	if err := json.Unmarshal(item, &place); err != nil {
		fmt.Println(err)
	} else {
		if place != *new(Place){
			places = append(places, place)
		}
	}
	return places
}