package main

import (
	"github.com/me/financial/config"
	"github.com/me/financial/router"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	router.InitializeRoutes()
}
