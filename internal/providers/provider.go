package providers

import (
	"github.com/ngocthanh06/authentication/cmd"
	"github.com/ngocthanh06/authentication/internal/database"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"github.com/ngocthanh06/authentication/internal/services"
)

var UserServ *services.UserService
var AuthServ *services.AuthService

func ConfigSetupProviders() {
	// execute command
	cmd.Execute()

	// repository
	UserRepo := repositories.NewUserRepository(database.DB)
	UserServ = services.NewUserService(UserRepo)
	AuthRepo := repositories.NewAuthRepository(database.DB)
	AuthServ = services.NewAuthService(AuthRepo)
}
