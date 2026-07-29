package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func RunMigrations(cfg DBConfig) error {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	))
	if err != nil {
		return fmt.Errorf("open db for migrations: %w", err)
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
