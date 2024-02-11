package db

import "fmt"

// DBInfo is an informational structure for database

type DBInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSL      bool
}

// String transforms db to connection string
func (db DBInfo) String() string {
	ssl := "disable"
	if db.SSL {
		ssl = "enable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s dbname=%s", db.Host, db.Port, db.User, db.Password, ssl, db.Name)
}
