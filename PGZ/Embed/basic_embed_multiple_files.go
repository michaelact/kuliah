package main

import (
	"database/sql"
	"io/fs"
	"embed"
	"time"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

//go:embed migrations/*.sql
var files embed.FS

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
						os.Getenv("MYSQL_USER"), 
						os.Getenv("MYSQL_PASSWORD"), 
						os.Getenv("MYSQL_HOST"), 
						os.Getenv("MYSQL_PORT"), 
						os.Getenv("MYSQL_DATABASE"))
	db := ConnectDB(dsn)
	defer db.Close()

	dirEntry, err := files.ReadDir("migrations")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntry {
		if !entry.IsDir() {
			content, err := fs.ReadFile(files, "migrations/" + entry.Name())
			if err != nil {
				panic(err)
			}

			db.Exec(string(content))
		}
	}
}
