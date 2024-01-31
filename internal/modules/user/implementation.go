package user

import (
	"delivery_api/cmd/configs"
	"delivery_api/internal/modules/user/controllers"
	"delivery_api/internal/modules/user/repository"
	usecases "delivery_api/internal/modules/user/use-cases"
	"delivery_api/pkg/auth/service"
	"delivery_api/pkg/infra/db"
)

type CreateUserController struct{}

func (impl CreateUserController) Create() controllers.CreateUserController {
	dbAdapter := db.InitDB()
	repository := repository.NewCreateUserRepo(dbAdapter)
	token := service.NewAuthJWTService(configs.GetSecretKey())
	usecase := usecases.NewCreateUserUseCase(repository, *token)
	return controllers.NewCreateUserController(usecase)
}


