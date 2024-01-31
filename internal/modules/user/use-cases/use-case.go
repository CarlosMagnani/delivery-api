package usecases

import (
	"delivery_api/internal/modules/user/models"
	repository "delivery_api/internal/modules/user/repository"
	"delivery_api/internal/responses"
	"delivery_api/pkg/auth/service"
)

type CreateUserUseCase struct {
	createUserRepo repository.CreateUserRepo
	authService service.AuthJWT
}

func NewCreateUserUseCase(createUserRepo repository.CreateUserRepo, authService service.AuthJWT) CreateUserUseCase {
	return CreateUserUseCase{createUserRepo, authService}
}

func (u *CreateUserUseCase) Execute( user models.User) error {
	user = models.User{ Name: user.Name, Description: user.Description, Password: user.Password, Status: user.Status}
	if err := u.createUserRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (u *CreateUserUseCase) FindUser(user models.UserLogin) (responses.LoginResponse, error) {

    if err := u.createUserRepo.FindUser(&user); err != nil {
        return responses.LoginResponse{}, err
    }
	token, err := u.authService.GenerateTokenJWT()
	if err != nil {
		return responses.LoginResponse{}, err
	}

	result := &responses.LoginResponse{Token: token, TokenType: "HS266"}
    return *result, nil
}