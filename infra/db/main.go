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

const UP_INDICATION = "_up.sql"
const DOWN_INDICATION = "_down.sql"

const UP = "up"
const DOWN = "down"

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("Please provide a direction argument (up or down)")
	}

	direction := args[1]

	if direction != "up" && direction != "down" {
		log.Fatal("Invalid argument (up or down)")
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := connectToDB()
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

		if !isUpFile(file.Name()) && direction == "up" {
			continue
		}
		if isUpFile(file.Name()) && direction == "down" {
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("[MIGRATION] executing file %s", file.Name())
		db.Exec(string(script))
	}
}

func connectToDB() (*sql.DB, error) {
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

func isUpFile(name string) bool {
	return name[len(name)-len(UP_INDICATION):] == UP_INDICATION
}
