package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to connect DB: %v", err)
	}
	return db, nil
}
