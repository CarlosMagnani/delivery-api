package repository

import (
	"delivery_api/internal/modules/user/models"
	"delivery_api/pkg/infra/db"
	"fmt"
)

type CreateUserRepo struct {
	dbAdapter db.GormAdapter
}

func NewCreateUserRepo( dbAdapter db.GormAdapter) CreateUserRepo {
	return CreateUserRepo{dbAdapter}
}


func(repository CreateUserRepo) Create(user models.User) error {
	result := repository.dbAdapter.GetDB().Table("public.tbl_user").Create(&user)
	fmt.Println(result.Error)
	if result.Error != nil {
		return result.Error
	}

	return nil
}