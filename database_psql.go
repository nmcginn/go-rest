package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func describe_tables(d postgres_db) (results []map[string]interface{}, err error) {
	conn_str := "user=" + d.Username + " password=" + d.Password + " dbname=" + d.Database + " host=" + d.Host + " port=" + d.Port + " sslmode=disable"
	db, err := sql.Open("postgres", conn_str)
	defer db.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return results, err
	}

	raw_query, err := Asset("data/describe.sql")
	if err != nil {
		return results, err
	}
	available_tables := string(raw_query[:])

	rows, err := db.Query(available_tables)
	defer rows.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return results, err
	}
	cols, _ := rows.Columns()

	for rows.Next() {
		columns := make([]interface{}, len(cols))
		column_pointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			column_pointers[i] = &columns[i]
		}
		err = rows.Scan(column_pointers...)
		if err != nil {
			return results, err
		}
		m := make(map[string]interface{})
		for i, col_name := range cols {
			val := column_pointers[i].(*interface{})
			m[col_name] = *val
		}
		results = append(results, m)
	}
	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return results, err
	}

	return results, nil
}

type postgres_db struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Schema   string
}
