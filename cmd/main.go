package main

import (
	"github.com/me/finance/config"
	"github.com/me/finance/config/logger"
	"github.com/me/finance/pkg/router"
)

func main() {
	logger.InitLogger()
	
	if err := config.Load(); err != nil {
		panic(err)
	}

	router.InitializeRoutes()
}
