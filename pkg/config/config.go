package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	CommandStart = "start"
	AdminID      = 2018211211
)

type Configuration struct { // TODO
	TOKEN             string
	DBName            string
	AdministratorsIDs map[string]int
}

func DBConnect() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "sqliteDB.db")
	// для хранения в ОЗУ имя БД ":memory:"
	return db, err
}
