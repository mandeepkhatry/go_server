package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	db *sql.DB
}

func NewClient(dir string) (*SQLite, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, os.ModePerm)
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/data.db", dir))
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	sqlite := &SQLite{db: db}
	if err := sqlite.createTables(); err != nil {
		return nil, err
	}
	return sqlite, nil
}

func (sq *SQLite) createTables() error {
	//Create Table if it doesn't exists
	tx, err := sq.db.Prepare(
		`CREATE TABLE IF NOT EXISTS data(
					id text UNIQUE NOT NULL PRIMARY KEY,
					file_name text NOT NULL,
					file_size int NOT NULL,
					content_type text NOT NULL
		);`)
	if err != nil {
		return err
	}
	_, err = tx.Exec()
	return err
}

func (sq *SQLite) CloseClient() error {
	return sq.db.Close()
}

func serialize(src interface{}) []byte {
	srcInBytes, _ := json.Marshal(src)
	return srcInBytes
}
