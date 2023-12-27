package controllers

import (
	"delivery_api/internal/modules/user/models"
	usecases "delivery_api/internal/modules/user/use-cases"
	"delivery_api/pkg/helpers"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateUserController struct {
	createUserUseCase usecases.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase usecases.CreateUserUseCase) CreateUserController {
	return CreateUserController{createUserUseCase}
}

func (controller CreateUserController) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	var userRequest *models.User

	teste := json.Unmarshal(body, &userRequest);
	fmt.Println(teste)
	if err != nil {
		helpers.ErrorResponse(w , http.StatusUnprocessableEntity, err)
	}
 

	if err = json.Unmarshal(body, &userRequest); err != nil {

		helpers.ErrorResponse(w, http.StatusBadRequest, err)
	}
     

	if err = controller.createUserUseCase.Execute(*userRequest); err != nil {
          helpers.ErrorResponse(w, http.StatusBadRequest, err)
		}

}