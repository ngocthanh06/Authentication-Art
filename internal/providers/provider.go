package providers

import (
	"github.com/ngocthanh06/authentication/cmd"
	"github.com/ngocthanh06/authentication/internal/database"
	"github.com/ngocthanh06/authentication/internal/repositories"
	"github.com/ngocthanh06/authentication/internal/services"
)

var UserServ services.UserServiceInterface
var AuthServ *services.AuthService
var UserRepo repositories.UserRepositoryInterface

func ConfigSetupProviders() {
	// execute command
	cmd.Execute()

	// repository
	UserRepo = repositories.NewUserRepository(database.DB)
	AuthRepo := repositories.NewAuthRepository(database.DB)
	UserServ = services.NewUserService(UserRepo)
	AuthServ = services.NewAuthService(AuthRepo, UserRepo)
}
