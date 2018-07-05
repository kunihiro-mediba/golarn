package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabase() {
	db, err := sql.Open("mysql", "test:test@/test")
	if err != nil {
		fmt.Println("connect DB failed", err)
		return
	}
	defer db.Close()

	// Execute query
	rows, err := db.Query("SELECT id,title,body FROM `article` ORDER BY `id`")
	if err != nil {
		fmt.Println("Query failed", err)
		return
	}
	defer rows.Close()

	// Read rows
	articles := make([]Article, 0)
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.ID, &article.Title, &article.Body)
		articles = append(articles, article)
	}

	// Generate JSON string
	buf, err := json.Marshal(articles)
	if err != nil {
		fmt.Println("JSON encode failed", err)
		return
	}

	fmt.Println(string(buf[:]))
}
