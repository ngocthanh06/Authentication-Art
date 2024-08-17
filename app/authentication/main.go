package main

import (
	"github.com/joho/godotenv"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/ngocthanh06/authentication/internal/database"
	"github.com/ngocthanh06/authentication/internal/providers"
	"github.com/ngocthanh06/authentication/internal/routes"
)

func init() {
	godotenv.Load(config.Env)
	config.InitEnvKey()
}

func main() {
	// connection database
	database.ConnectionDatabase()

	// ConfigSetupProviders
	providers.ConfigSetupProviders()

	//create routes
	routes.CreateRoutes()

}
