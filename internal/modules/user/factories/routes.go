package factories

import (
	"delivery_api/internal/modules/user/routes"
	"delivery_api/pkg/types"
)

func BuildRoutes(httpAdapter types.HttpAdapter){
	for _, route := range routes.UserRoutes {
		httpAdapter.CreateRoute(types.Route{
			Uri: route.Uri,
			Method: route.Method,
			Function: route.Function,
		})
	}
}