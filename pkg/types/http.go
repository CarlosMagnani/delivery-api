package types

type HttpAdapter interface {
	CreateRoute(route Route)
}