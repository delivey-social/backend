package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"comida.app/src/infra/environment"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	host := environment.DB_HOST
	port := environment.DB_PORT
	user := environment.DB_USER
	password := environment.DB_PASS
	dbname := environment.DB_NAME

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
		return nil, fmt.Errorf("[DB] Error connecting: %w", err)
	}

	slog.Info("[DB] Connected. Pinging...")

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("[DB] Error pinging: %w", err)

	}

	slog.Info("[DB] Connected successfully")

	return db, nil
}
