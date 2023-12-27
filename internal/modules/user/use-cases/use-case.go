package usecases

import (
	"delivery_api/internal/modules/user/models"
	repository "delivery_api/internal/modules/user/repository"
	"fmt"
)

type CreateUserUseCase struct {
	createUserRepo repository.CreateUserRepo
}

func NewCreateUserUseCase(createUserRepo repository.CreateUserRepo) CreateUserUseCase {
	return CreateUserUseCase{createUserRepo}
}

func (u *CreateUserUseCase) Execute( user models.User) error {
	user = models.User{ Name: user.Name, Description: user.Description, Password: user.Password, Status: user.Status}
	fmt.Println(user)
	if err := u.createUserRepo.Create(user); err != nil {
		return err
	}

	return nil
}