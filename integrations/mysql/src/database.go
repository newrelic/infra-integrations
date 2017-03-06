package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type dataSource interface {
	close()
	query(string) (map[string]interface{}, error)
}

type database struct {
	source *sql.DB
}

func openDB(dsn string) (dataSource, error) {
	source, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db := database{
		source: source,
	}
	return &db, nil
}

func (db *database) close() {
	db.source.Close()
}

func (db *database) query(query string) (map[string]interface{}, error) {
	rows, err := db.source.Query(query)
	if err != nil {
		return nil, err
	}

	rawData := make(map[string]interface{})
	var name, value string

	for rows.Next() {
		if err = rows.Scan(&name, &value); err != nil {
			return nil, err
		}
		rawData[name] = asValue(value)
	}
	return rawData, nil

}
