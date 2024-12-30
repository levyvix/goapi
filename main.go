package main

import (
	"github.com/levyvix/goapi/config"
	"github.com/levyvix/goapi/router"
)

var (
	logger *config.Logger
)

func main() {

	// init logger
	logger = config.GetLogger("")

	// Init db
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing db: %v", err)
		return
	}

	router.Initialize()
}
