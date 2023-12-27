package main

import (
	"delivery_api/cmd/configs"
	"delivery_api/internal/modules/user/factories"
	"delivery_api/pkg/adapters"
	"delivery_api/pkg/types"
	"fmt"
	"log"
	"net/http"
)

func main() {

	err := configs.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	runMessage := fmt.Sprintf("🚀 task service is running at http://localhost:%s", "4000")
	httpAdapter := adapters.NewChiAdapter()
	httpAdapter.CreateRoute(types.Route{
		Uri:    "/",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(runMessage))
		},
	})

	factories.BuildRoutes(httpAdapter)
	fmt.Println(runMessage)
	for {
		http.ListenAndServe(fmt.Sprintf(":%d", 4000), httpAdapter.GetRouter())
	}
}