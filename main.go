package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/andev58/ira-pkgsearch/backend"
	"github.com/andev58/ira-pkgsearch/backend/constants"
)

const DEFAULT_STAGE = constants.STAGE_DEV
const DEFAULT_PORT = 9870

func main() {
	defaultPkgDir, _ := filepath.Abs("./pkgs")

	addr := flag.String("addr", fmt.Sprintf(":%d", DEFAULT_PORT), "HTTPS network address")
	certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
	keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	pkgDir := flag.String("pkgdir", defaultPkgDir, "Directory where IRA packages are stored")
	flag.Parse()

	backend.Run(DEFAULT_STAGE, *addr, *certFile, *keyFile, *pkgDir)
}
