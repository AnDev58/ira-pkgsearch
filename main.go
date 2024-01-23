// Package main starts server with a special settings
package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/andev58/ira-pkgsearch/backend"
	"github.com/andev58/ira-pkgsearch/backend/constants"
)

const defaultStage = constants.StageDev
const defaultPort = 9870

func main() {
	defaultPkgDir, _ := filepath.Abs("./pkgs")

	addr := flag.String("addr", fmt.Sprintf(":%d", defaultPort), "HTTPS network address")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	pkgDir := flag.String("pkgdir", defaultPkgDir, "Directory where IRA packages are stored")
	flag.Parse()

	backend.Run(defaultStage, *addr, *certFile, *keyFile, *pkgDir)
}
