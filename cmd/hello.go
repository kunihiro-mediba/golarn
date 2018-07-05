package main

import (
	"fmt"

	"github.com/kunihiro-mediba/golarn/storage"
)

func main() {
	fmt.Println("Hello")

	// Test storage
	storage.TestDatabase()

	// Test memcached
	storage.TestMemcached()

	// Test JSON
	storage.TestJson()

	// Test XML
	storage.TestXml()
}
