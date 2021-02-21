package infra

import (
	"database/sql"
	"fmt"
	"os"
)

func initDb() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", os.Getenv("dsn"))

	if err != nil {
		return nil, fmt.Errorf("failed to connect mysql %w", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect mysql %w", err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	return dbMap, nil
}
