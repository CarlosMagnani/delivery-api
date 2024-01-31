package main

import (
	"delivery_api/cmd/configs"
	"delivery_api/internal/modules/user/routes"
	"delivery_api/pkg/adapters"
	"delivery_api/pkg/auth/middleware"
	"delivery_api/pkg/auth/service"
	"delivery_api/pkg/factories"
	"delivery_api/pkg/types"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	runMessage := fmt.Sprintf("🚀 delivery api is running at http://localhost:%s", "4000")
	httpAdapter := adapters.NewGinAdapter()
	secretKey := configs.GetSecretKey()
	authService := service.NewAuthJWTService(secretKey)
	factories.BuildRoutes(httpAdapter, routes.UserRoutes)
	httpAdapter.GetRouter().Use(middleware.AuthMiddleware(authService))
	httpAdapter.CreateRoute(types.Route{
		Uri:    "/",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *gin.Context) {
			w.Write([]byte(runMessage))
		},
	})

	
	fmt.Println(runMessage)
	for {
		http.ListenAndServe(fmt.Sprintf(":%d", 4000), httpAdapter.GetRouter())
	}
}