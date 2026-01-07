package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	_, err := Connect()
	if err != nil {
		fmt.Println(err)
	}
}

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("HOST", host)

	psqlInfo := fmt.Sprintf(`
        host=%s
        port=%s
        user=%s
        password=%s
        dbname=%s
        sslmode=disable
    `, host, port, user, password, dbname)

	slog.Info("[DB] Connecting...")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	slog.Info("[DB] Connected. Pinging...")

	err = db.Ping()
	if err != nil {
		return nil, err

	}

	slog.Info("[DB] Connected successfully")

	return db, nil
}
