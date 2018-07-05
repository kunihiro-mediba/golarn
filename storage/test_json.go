package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JSONData struct {
	Result   string    `json:"result"`
	Articles []Article `json:"articles"`
}

func TestJson() {
	// Read file
	bytes, err := ioutil.ReadFile("test/testdata/test.json")
	if err != nil {
		fmt.Println("Read file failed", err)
		return
	}

	// Decode
	data := JSONData{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Println("JSON parse failed", err)
		return
	}

	// Add item
	article := Article{
		ID: 101,
		Title: "added",
		Body: "Test added testdata",
	}
	data.Articles = append(data.Articles, article)

	// Encode
	encoded, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encode failed", err)
		return
	}
	fmt.Println(string(encoded))
}
