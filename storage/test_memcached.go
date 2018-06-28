package storage

import (
	"fmt"
	"time"

	"github.com/rainycape/memcache"
)

func TestMemcached() {

	client, err := memcache.New("localhost:11211")
	if err != nil {
		fmt.Println("Connect memcached failed", err)
		return
	}
	defer client.Close()

	key := "lastAccess"

	// Get cache
	item, err := client.Get(key)
	if err != nil {
		fmt.Println("Memcached Get failed", err)
	} else {
		fmt.Printf("Memcached get key='%s'\n", key)
		fmt.Printf("%s = %s\n", key, string(item.Value))
	}

	// Update cache
	value := time.Now().Format("2006-01-02 15:04:05")
	err = client.Set(&memcache.Item{Key: key, Value: []byte(value)})
	if err != nil {
		fmt.Println("Memcache set failed", err)
		return
	}
	fmt.Println("Memcached set success")

	// Get cache (retry)
	item, err = client.Get(key)
	if err != nil {
		fmt.Println("Memcached get failed", err)
	} else {
		fmt.Printf("Memcached get key='%s'\n", key)
		fmt.Printf("%s = %s\n", key, string(item.Value))
	}
}
