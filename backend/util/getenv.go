package util

import (
	"log"
	"os"
	"strconv"

	"github.com/andev58/ira-pkgsearch/backend/constants"
)

func GetServerStage(defaultStage int) int {
	stage := defaultStage
	if mode, ok := os.LookupEnv("MODE"); ok {
		switch mode {
		case "DEVELOPMENT":
			stage = constants.STAGE_DEV
		case "TEST":
			stage = constants.STAGE_TEST
		case "PRODUCTION":
			stage = constants.STAGE_PROD
		default:
			log.Fatal("incorrect MODE value")
		}
	}
	return stage
}

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
