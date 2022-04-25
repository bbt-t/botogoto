package botogoto_mainbody

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type UserSchema struct {
	Time       time.Time
	Name       string
	TelegramId int64
}

func CreateSchema(db *sql.DB) {
	query, _ := db.Prepare(`
			CREATE TABLE IF NOT EXISTS "user_info" (
			    "time" DATE NOT NULL,
			    "name" VARCHAR(32),
			    "telegram_id" INTEGER PRIMARY KEY AUTOINCREMENT
			);
	`)
	if _, err := query.Exec(); err != nil {
		log.Printf("Create DB schema error: %s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func Add(db *sql.DB, user UserSchema) {
	query, _ := db.Prepare("INSERT INTO user_info (time, name, telegram_id) VALUES (?, ?, ?)")

	if _, err := query.Exec(user.Time, user.Name, user.TelegramId); err != nil {
		log.Printf("Create DB schema error: %s", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}
