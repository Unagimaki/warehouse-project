package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func RunMigrations(cfg DBConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open(cfg.Driver, dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				break
			}
			db.Close()
		}

		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	projectRoot := "."
	if wd, err := os.Getwd(); err == nil {
		projectRoot = wd
	}

	migrationPath := filepath.Clean(filepath.Join(projectRoot, "migrations", "001_create_products_table.sql"))
	if _, err := os.Stat(migrationPath); err != nil {
		return fmt.Errorf("migration file not found: %s: %w", migrationPath, err)
	}

	sqlBytes, err := os.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("read migration file: %w", err)
	}

	if _, err := db.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}
