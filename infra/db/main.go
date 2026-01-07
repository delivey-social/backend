package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir("migration")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join("migration", file.Name())
		script, err := os.ReadFile(path)

		if err != nil {
			log.Fatal(err)
		}

		db.Exec(string(script))
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
