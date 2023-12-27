package routes

import (
	factories "delivery_api/internal/modules/user/factories/controllers"
	"delivery_api/pkg/types"
	"net/http"
)

var (
	createUserController = factories.CreateUserControllerFactory{}.Create()
)

var UserRoutes = []types.Route{
	{
		Uri: "/user",
		Method: http.MethodPost,
		Function: createUserController.Handle,
	},
}