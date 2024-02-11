// Package main starts server with a special settings
package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/andev58/ira-pkgsearch/backend"
	"github.com/andev58/ira-pkgsearch/backend/constants"
	"github.com/andev58/ira-pkgsearch/backend/db"
)

const defaultStage = constants.StageDev
const defaultPort = 9870
const defaultPortHTTP = 7890

var DBInfo = db.DBInfo{
	Name:     "iraOrigin",
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "root",
	SSL:      false,
}

func main() {
	defaultPkgDir, _ := filepath.Abs("./pkgs")

	addrTLS := flag.String("addr", fmt.Sprintf(":%d", defaultPort), "HTTPS network address")
	useHTTPS := flag.Bool("tls", true, "Whether use HTTPS or not")
	useHTTP := flag.Bool("http", false, "Whether use HTTP or not")
	addrHTTP := flag.String("legacy-addr", fmt.Sprintf(":%d", defaultPortHTTP), "Legacy HTTP network address")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	pkgDir := flag.String("pkgdir", defaultPkgDir, "Directory where IRA packages are stored")

	dbName := flag.String("db-name", DBInfo.Name, "Name of PostgreSQL database")
	dbHost := flag.String("db-host", DBInfo.Host, "PostgreSQL host")
	dbPort := flag.Int("db-port", DBInfo.Port, "PostgreSQL port")
	dbUser := flag.String("db-user", DBInfo.User, "PostgreSQL user")
	dbUserPwd := flag.String("db-userpwd", DBInfo.Password, "PostgreSQL user's password")
	dbUseSSL := flag.Bool("db-ssl", DBInfo.SSL, "If set, server will connect to database using SSL")

	flag.Parse()

	DBInfo = db.DBInfo{Name: *dbName, Host: *dbHost, Port: *dbPort, User: *dbUser, Password: *dbUserPwd, SSL: *dbUseSSL}
	backend.Run(defaultStage, *useHTTP, *addrHTTP, *useHTTPS, *addrTLS, *certFile, *keyFile, *pkgDir, DBInfo)
}
