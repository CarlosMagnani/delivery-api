package factories

import (
	"delivery_api/internal/modules/user/repository"
	usecases "delivery_api/internal/modules/user/use-cases"
	"delivery_api/pkg/infra/db"
)

type CreateUserUseCaseFactory struct{}

func (fac CreateUserUseCaseFactory) Create() usecases.CreateUserUseCase {
	dbAdapter := db.InitDB()
	repository := repository.NewCreateUserRepo(dbAdapter)
	return usecases.NewCreateUserUseCase(repository)
}