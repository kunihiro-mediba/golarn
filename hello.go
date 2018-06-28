package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kunihiro-mediba/golarn/storage"
)

func main() {
	fmt.Println("Hello")

	// Test storage
	storage.TestDatabase()

	// Test memcached
	storage.TestMemcached()
}
