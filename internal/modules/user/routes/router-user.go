package routes

import (
	"delivery_api/internal/modules/user"
	"delivery_api/pkg/types"
	"net/http"
)

var (
	createUserController = user.CreateUserController{}.Create()
)

var UserRoutes = []types.Route{
	{
		Uri: "/user",
		Method: http.MethodPost,
		Function: createUserController.Handle,
	},
	{
		Uri: "/login",
		Method: http.MethodPost,
		Function: createUserController.SignIn,
	},
}