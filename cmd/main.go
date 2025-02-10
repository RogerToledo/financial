package main

import (
	"github.com/me/financial/config"
	"github.com/me/financial/config/logger"
	"github.com/me/financial/pkg/router"
)

func main() {
	logger.InitLogger()
	
	if err := config.Load(); err != nil {
		panic(err)
	}

	router.InitializeRoutes()
}
