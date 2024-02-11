// Package db specifies database connections
package db

import (
	_ "github.com/jackc/pgx/v5/stdlib" // DB driver
	"github.com/jmoiron/sqlx"
)

// Connect sets connection to database by specified DBInfo
func Connect(info DBInfo) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", info.String())
}
