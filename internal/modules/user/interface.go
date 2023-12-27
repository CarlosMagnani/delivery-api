package user

import "delivery_api/internal/modules/user/models"

type UserRepo interface {
	Create(Model models.User) error
}