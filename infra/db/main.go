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

type Direction string

const UP = "up"
const DOWN = "down"

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatal("Please provide a direction argument (up or down)")
	}

	direction, ok := newDirection(args[1])
	if !ok {
		log.Fatalf("Invalid argument: Expected 'up' or 'down' got %s\n", direction)
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
		if err != nil {
			log.Fatalf("Error reading file %s", file.Name())
		}

		fileDirection, ok := getFileDirection(file.Name())
		if !ok {
			log.Fatalf("[MIGRATION] File with invalid direction detected %s\n", file.Name())
		}

		if fileDirection != direction {
			continue
		}

		log.Printf("[MIGRATION] executing file %s", file.Name())
		_, err = db.Exec(string(script))
		if err != nil {
			log.Fatalf("[MIGRATION] Error while executing %s: %s\n", file.Name(), err.Error())
		}
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

func newDirection(try string) (Direction, bool) {
	return Direction(try), try == UP || try == DOWN
}

func getFileDirection(fileName string) (Direction, bool) {
	isUp := fileName[len(fileName)-len(UP_INDICATION):] == UP_INDICATION
	isDown := fileName[len(fileName)-len(DOWN_INDICATION):] == DOWN_INDICATION

	if isUp {
		return UP, true
	}

	if isDown {
		return DOWN, true
	}

	return "", false
}
