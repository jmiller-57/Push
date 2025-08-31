package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			roomname TEXT NOT NULL UNIQUE,
			creator_id INTEGER,
			FOREIGN KEY (creator_id) REFERENCES users(id)
		);
		CREATE TABLE IF NOT EXISTS room_members (
			room_id INTEGER,
			user_id INTERGER,
			PRIMARY KEY (room_id, user_id),
			FOREIGN KEY (room_id) REFERENCES rooms(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`)
	if err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}
	return db
}
