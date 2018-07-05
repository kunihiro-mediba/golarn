package storage

import (
	"fmt"
	"io/ioutil"

	"encoding/xml"
)

type XMLData struct {
	TagName  xml.Name  `xml:"data"`
	Result   string    `xml:"result"`
	Articles []Article `xml:"articles"`
}

func TestXml() {
	// Read file
	bytes, err := ioutil.ReadFile("test/testdata/test.xml")
	if err != nil {
		fmt.Println("Read file failed", err)
		return
	}

	// Decode
	data := XMLData{}
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("XML parse failed", err)
		return
	}
	fmt.Println(data)

	// Add item
	article := Article{
		ID: 101,
		Title: "added",
		Body: "Test added testdata",
	}
	data.Articles = append(data.Articles, article)

	// Encode
	encoded, err := xml.Marshal(data)
	if err != nil {
		fmt.Println("XML encode failed", err)
		return
	}
	fmt.Println(string(encoded))
}
