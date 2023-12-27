package adapters

import (
	"delivery_api/pkg/types"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiAdapter struct {
	router *chi.Mux
}

func NewChiAdapter() chiAdapter {
	return chiAdapter{chi.NewRouter()}
}

func (adapter chiAdapter) CreateRoute(route types.Route) {
	switch route.Method {
	case http.MethodGet:
		adapter.router.Get(route.Uri, route.Function)

	case http.MethodPost:
		adapter.router.Post(route.Uri, route.Function)

	case http.MethodPut:
		adapter.router.Put(route.Uri, route.Function)

	case http.MethodDelete:
		adapter.router.Delete(route.Uri, route.Function)

	default:
		log.Fatal("Invalid http method!")
	}
}

func (adapter chiAdapter) GetRouter() *chi.Mux {
	return adapter.router
}