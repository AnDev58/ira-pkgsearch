package main

import (
	"github.com/andev58/ira-pkgsearch/backend"
	"github.com/andev58/ira-pkgsearch/backend/constants"
)

const DEFAULT_STAGE = constants.STAGE_DEV
const DEFAULT_PORT = 9870

func main() {
	backend.Run(DEFAULT_PORT, DEFAULT_STAGE)
}
