// Package db specifies database connections
package db

import (
	"database/sql"

	_ "github.com/lib/pq" // DB driver
)

// Connect sets connection to database by specified DBInfo
func Connect(info DBInfo) (*sql.DB, error) {
	return sql.Open("postgres", info.String())
}
