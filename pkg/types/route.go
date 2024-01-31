package types

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Uri    string
	Method string
	Function func(w http.ResponseWriter, r *gin.Context)
}