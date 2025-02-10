package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/me/financial/config"
	"github.com/me/financial/pkg/db"
	"github.com/me/financial/pkg/repository"
	"github.com/sagikazarmark/slog-shim"
)

func InitializeRoutes() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	defer func () {
		if err := db.Close(); err != nil {
			log.Printf("error trying close database: %v", err)
		}
	}()	

	rep := repository.NewRepository(db)

	mux := http.NewServeMux()
	
	PersonRoutes(mux, rep)
	CreditCardRoutes(mux, rep)
	PaymentTypeRoutes(mux, rep)
	PurchaseTypeRoutes(mux, rep)
	PurchaseRoutes(mux, rep)

	slog.Info("Server running on port %s", config.ServerPort())
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort()), mux)
}