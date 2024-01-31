package adapters

import (
	"delivery_api/pkg/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginAdapter struct {
	router *gin.Engine
}

func NewGinAdapter() ginAdapter {
	return ginAdapter{gin.Default()}
}
	
func (adapter ginAdapter) CreateRoute(route types.Route) {
	switch route.Method {
	case http.MethodGet:
		adapter.router.GET(route.Uri,  ginAdapterHandler(route.Function))

	case http.MethodPost:
		adapter.router.POST(route.Uri,  ginAdapterHandler(route.Function))

	case http.MethodPut:
		adapter.router.PUT(route.Uri, ginAdapterHandler(route.Function))

	case http.MethodDelete:
		adapter.router.DELETE(route.Uri,  ginAdapterHandler(route.Function))

	default:
		log.Fatal("Invalid http method!")
	}
}

func (adapter ginAdapter) GetRouter() *gin.Engine {
	return adapter.router
}
func ginAdapterHandler(handlerFunc func(w http.ResponseWriter, r *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handlerFunc(c.Writer, c)
	}
}