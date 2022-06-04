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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
						os.Getenv("MYSQL_USER"), 
						os.Getenv("MYSQL_PASSWORD"), 
						os.Getenv("MYSQL_HOST"), 
						os.Getenv("MYSQL_PORT"), 
						os.Getenv("MYSQL_DATABASE"))
	db := ConnectDB(dsn)
	defer db.Close()

	// Gunakan .QueryContext() untuk eksekusi query yang mengembalikan data
	ctx := context.Background()
	script := "SELECT id, firstname, lastname, email, reg_date from MyGuests;"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var firstname, lastname, email string
		var regDate time.Time

		// Mengikuti urutan saat melakukan Select Column
		err := rows.Scan(&id, &firstname, &lastname, &email, &regDate)
		if err != nil {
			panic(err)
		}

		fmt.Println("Firstname:", firstname)
	}

	fmt.Println("Success select new data.")
}
