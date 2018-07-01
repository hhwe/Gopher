package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
	title    string
	finished bool
	created  time.Time
}

var db, err = sql.Open("mysql", "root:root@/todo?charset=utf8")
var rd = redis.NewClient(&redis.Options{DB: 0})

func getAll(query *sql.Rows) []map[string]string {
	column, _ := query.Columns()
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}
	results := make([]map[string]string, 0)
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			fmt.Println(err)
			return nil
		}
		row := make(map[string]string)
		for k, v := range values {
			key := column[k]
			row[key] = string(v)
		}
		results = append(results, row)
	}
	return results
}
