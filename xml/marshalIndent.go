package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type TNote struct {
	Lang    string `xml:"lang,attr"`
	Content string `xml:",innerxml"`
}

type TFile struct {
	XMLName  struct{} `xml:"file"`
	FileName string   `xml:"name,attr"`
	Size     string   `xml:"size,attr"`
}

type Map map[string]string

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name,attr"`
	Age     string   `xml:"age,attr"`
}

type Release struct {
	XMLName   struct{} `xml:"release"`  // 使用此字段tag上定义的属性进行解析
	Comment   string   `xml:",comment"` // xml中的注释
	Version   string   `xml:"version,attr"`
	TimeStamp string   `xml:",attr"`
	Lang      string   `xml:"-"`
	Skin      string   `xml:",chardata"`
	Site      string   `xml:",omitempty"`
	File      []TFile  `xml:",innerxml"`
	CnNotes   TNote    `xml:"cnnote"`
	EnNotes   TNote    `xml:"ennote"`
}

func testMarshal() {
	release := Release{
		Version:   "1.0.0.0",
		TimeStamp: time.Now().String(),
		Lang:      "zh-cn",
		File: []TFile{TFile{FileName: "/deploy/package_1.zip", Size: "50"},
			TFile{FileName: "/deploy/package_2.zip", Size: "60"},
		},
		Skin:    "blue",
		Comment: "this is a test for xml parser",
	}

	v, err := xml.MarshalIndent(release, "", "   ")
	if err != nil {
		fmt.Println("marshal xml value error, error msg:%s", err.Error())
	}
	fmt.Println("marshal xml value:\n", string(v))

	fmt.Println("-----------xml parse-------------")
	m := Map{
		"Name": "Chris",
		"Age":  "25",
	}

	p := make([]Person, 0)

	for k, v := range m {
		p = append(p, Person{Name: k, Age: v})
	}

	v, err = xml.MarshalIndent(p, "", "   ")
	if err != nil {
		fmt.Println("marshal xml value error, error msg:%s", err.Error())
	}
	fmt.Println("convert map to xml value\n", string(v))
}
