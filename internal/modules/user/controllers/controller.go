package controllers

import (
	"delivery_api/internal/modules/user/models"
	usecases "delivery_api/internal/modules/user/use-cases"
	"delivery_api/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserController struct {
	createUserUseCase usecases.CreateUserUseCase
}

func NewCreateUserController(createUserUseCase usecases.CreateUserUseCase) CreateUserController {
	return CreateUserController{createUserUseCase}
}

func (controller CreateUserController) Handle(w http.ResponseWriter, r *gin.Context) {
	var userRequest models.User
	if  r.Bind(&userRequest) != nil {
		helpers.ErrorResponse(w , http.StatusUnprocessableEntity, "Error on process your entity")
	}
 
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10); 
	if err != nil {
		helpers.ErrorResponse(w, http.StatusUnprocessableEntity, "Error on process your entity aqui ó")
	}

	userRequest.Password = string(hash)
	if err = controller.createUserUseCase.Execute(userRequest); err != nil {
          helpers.ErrorResponse(w, http.StatusBadRequest, "Bad request error on repository")
		}

}


func (controller CreateUserController) SignIn(w http.ResponseWriter, r *gin.Context){
	var bodyUser models.UserLogin;
	if r.Bind(&bodyUser) != nil {
		helpers.ErrorResponse(w , http.StatusUnprocessableEntity, "Error on process your entity")
	};
	fmt.Println("CONTROLLER", bodyUser)
	token, err := controller.createUserUseCase.FindUser(bodyUser);
	if err != nil {
		helpers.ErrorResponse(w, http.StatusBadRequest, "Bad request error on repository")
		return
	};	

	helpers.JsonResponse(w, http.StatusOK, token)
};