package main

import (
	"faba_traning_project/cmd/handlers/routes"
	"faba_traning_project/internal/config"
	"faba_traning_project/internal/databases"
	"faba_traning_project/internal/databases/order"
	"faba_traning_project/internal/databases/order_item"
	"faba_traning_project/internal/databases/product"
	"faba_traning_project/internal/databases/user"
	"faba_traning_project/internal/services"
	"faba_traning_project/utils"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start the database")
	}

	// database store
	dbStore := databases.DBStore{
		User:      user.NewManagement(db),
		Product:   product.NewManagement(db),
		Order:     order.NewManagement(db),
		OrderItem: order_item.NewManagement(db),
		DBconn:    db,
	}

	// initialize all services inside a service container
	serviceContainer := services.Container{
		User:    services.NewUser(dbStore),
		Product: services.NewProduct(dbStore),
		Order:   services.NewOrder(dbStore),
	}

	r := routes.Route(serviceContainer)
	port := utils.GetWithDefault("API_PORT", "8000")
	log.Info().Msg("Faba Training service start on port " + port)
	http.ListenAndServe(":"+port, r)

}
