package infra

import (
	"database/sql"
	"fmt"

	"github.com/YoshitakaSS/go_auth/config"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// InitDB connect MySQL and return gorp mapping object
func InitDB() (*gorp.DbMap, error) {
	db, err := sql.Open("mysql", config.DBDSN())

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
