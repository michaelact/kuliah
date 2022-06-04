package main

import (
	"database/sql"
	"context"
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

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
						os.Getenv("MYSQL_USER"), 
						os.Getenv("MYSQL_PASSWORD"), 
						os.Getenv("MYSQL_HOST"), 
						os.Getenv("MYSQL_PORT"), 
						os.Getenv("MYSQL_DATABASE"))
	db := ConnectDB(dsn)
	defer db.Close()

	// Start database transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Prepare database statement
	ctx := context.Background()
	script := "INSERT INTO MyGuests(firstname, lastname, email) VALUES (?, ?, ?);"
	stmt, err := tx.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	for i := 0; i < 10; i++ {
		firstname := "Michael"
		lastname := "Act"
		email := "example@gmail.com"

		result, err := stmt.ExecContext(ctx, firstname, lastname, email)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Guest ID:", id)
	}

	// tx.Commit() // Complete Database Operation
	tx.Rollback() // Cancel Database Operation
}
