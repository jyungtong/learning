package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func openDB() *sql.DB {
	db, err := sql.Open("sqlite", "users.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to users.db")
	return db
}

func ddl(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func insertDemo(db *sql.DB) {
	res, err := db.Exec(`
		INSERT INTO users (name, age)
		VALUES (?, ?)`,
		"John", 30,
	)

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted id=%d\n", id)

	rows := []struct {
		name string
		age  int
	}{
		{"Bob", 25},
		{"Carol", 35},
		{"Dave", 28},
	}

	for _, r := range rows {
		res, err := db.Exec(
			"INSERT INTO users (name, age) VALUES (?, ?)",
			r.name, r.age,
		)
		if err != nil {
			log.Fatal(err)
		}

		id, _ := res.LastInsertId()
		fmt.Printf("inserted id=%d name=%s age=%d\n", id, r.name, r.age)
	}
}

func queryDemo(db *sql.DB) {
	var (
		id   int
		name string
		age  int
	)

	err := db.QueryRow(
		"SELECT id, name, age FROM users WHERE name = ?", "Alice",
	).Scan(&id, &name, &age)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("not found")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("QueryRow: id=%d name=%s age=%d\n", id, name, age)
	}

	rows, err := db.Query("SELECT id, name, age FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Query all users:")
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  id=%d name=%s age=%d\n", id, name, age)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func updateDeleteDemo(db *sql.DB) {
	res, err := db.Exec(
		"UPDATE users SET age = ? WHERE name = ?",
		33, "John", 
	)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("updated %d rows\n", n)

	res, err = db.Exec(
		"DELETE FROM users WHERE name = ?",
		"Dave",
	)
	if err != nil {
		log.Fatal(err)
	}
	n, _ = res.RowsAffected()
	fmt.Printf("deleted %d rows\n", n)
}

func main() {
	db := openDB()
	defer db.Close()

	ddl(db)
	insertDemo(db)
	queryDemo(db)
	updateDeleteDemo(db)
}
