package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateSchema(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Error creating migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migrations: %v", err)
	}

	log.Println("Successfully applied migrations")
}
