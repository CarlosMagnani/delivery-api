package factories

import (
	"delivery_api/pkg/types"
)

func BuildRoutes(httpAdapter types.HttpAdapter, routeSlices ...[]types.Route) {
	for _, routes := range routeSlices {
		for _, route := range routes {
			httpAdapter.CreateRoute(types.Route{
				Uri:      route.Uri,
				Method:   route.Method,
				Function: route.Function,
			})
		}
	}
}