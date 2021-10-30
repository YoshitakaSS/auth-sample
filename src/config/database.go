package config

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// InitDB connect MySQL and return gorp mapping object
func InitDB() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", DBDSN())

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
