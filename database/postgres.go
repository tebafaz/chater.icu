package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

// Postgres a
var postgres *sqlx.DB

// InitPostgres a
func InitPostgres() error {
	var err error
	postgres, err = sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSL_MODE")))
	if err != nil {
		return err
	}
	fmt.Printf("connection to %s created\n", postgres.DriverName())
	return err
}

// ClosePostgres a
func ClosePostgres() {
	err := postgres.Close()
	if err != nil {
		panic("postgres refused to close")
	}
}
