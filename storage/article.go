package storage

import (
	"encoding/xml"
)

type Article struct {
	TagName xml.Name `json:"-"     xml:"article"`
	ID      int      `json:"id"    xml:"id"`
	Title   string   `json:"title" xml:"title"`
	Body    string   `json:"body"  xml:"body"`
}
