package main

import (
	"github.com/joho/godotenv"
	"github.com/ngocthanh06/authentication/internal/config"
	"github.com/ngocthanh06/authentication/internal/database"
	"github.com/ngocthanh06/authentication/internal/providers"
	"github.com/ngocthanh06/authentication/internal/routes"
	"github.com/ngocthanh06/authentication/internal/utils"
)

func init() {
	godotenv.Load(utils.Env)
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
