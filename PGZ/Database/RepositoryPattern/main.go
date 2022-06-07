package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"context"
	"time"
	"fmt"
	"os"

	"repository_pattern/repository"
)

func connectdb(dsn string) *sql.DB {
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

	ctx := context.Background()

	guestRepository := repository.NewGuestRepository(connectdb(dsn))
	fmt.Println(guestRepository.FindById(ctx, 1))
}
