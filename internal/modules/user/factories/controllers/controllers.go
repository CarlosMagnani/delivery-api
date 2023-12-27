package factories

import (
	"delivery_api/internal/modules/user/controllers"
	factories "delivery_api/internal/modules/user/factories/use-cases"
)

type CreateUserControllerFactory struct{}

func (fac CreateUserControllerFactory) Create() controllers.CreateUserController {
	createUserUseCase := factories.CreateUserUseCaseFactory{}
	return controllers.NewCreateUserController(createUserUseCase.Create())
}