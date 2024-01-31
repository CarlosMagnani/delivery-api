package repository

import (
	"delivery_api/internal/modules/user/models"
	"delivery_api/pkg/infra/db"


































































	"golang.org/x/crypto/bcrypt"
)

type CreateUserRepo struct {
	dbAdapter db.GormAdapter
}

func NewCreateUserRepo( dbAdapter db.GormAdapter) CreateUserRepo {
	return CreateUserRepo{dbAdapter}
}


func(repository CreateUserRepo) Create(user models.User) error {
	result := repository.dbAdapter.GetDB().Table("public.tbl_user").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository CreateUserRepo) FindUser(user *models.UserLogin) error {
	var storedHashedPassword string
    result := repository.dbAdapter.GetDB().Table("public.tbl_user").Where("email = ?", user.Email).Select("password").First(&storedHashedPassword)
    if result.Error != nil {
        return result.Error
    }

	err := bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(user.Password))
	if err != nil {
		return err
	}

    return nil
}