// Package controllers set a default controllers
package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

// VersionPageHandler print version of each used program in browser
func VersionPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go version: %s\n", runtime.Version())

	node, err := exec.Command("node", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "NodeJS version: %s", node)

	yarn, err := exec.Command("npx", "yarn", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Yarn version: %s", yarn)

	vite, err := exec.Command("npx", "vite", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Vite version: %s", vite)
}
