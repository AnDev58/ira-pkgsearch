// Package util provides special features for a programmer
package util

import (
	"log"
	"os"
	"strconv"

	"github.com/andev58/ira-pkgsearch/backend/constants"
)

// GetServerStage gets server stage from environment or defaultStage if not present.
// Return a Stage constant
func GetServerStage(defaultStage int) int {
	stage := defaultStage
	if mode, ok := os.LookupEnv("MODE"); ok {
		switch mode {
		case "DEVELOPMENT":
			stage = constants.StageDev
		case "TEST":
			stage = constants.StageTest
		case "PRODUCTION":
			stage = constants.StageProd
		default:
			log.Fatal("incorrect MODE value")
		}
	}
	return stage
}

// GetPort gets port from environment or defaultPort otherway
func GetPort(defaultPort int) int {
	portNumber := defaultPort

	if portEnv, ok := os.LookupEnv("PORT"); ok {
		var err error
		portNumber, err = strconv.Atoi(portEnv)
		if err != nil {
			log.Fatal(err)
		}
	}
	return portNumber
}
