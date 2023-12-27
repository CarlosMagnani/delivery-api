package types

import "net/http"

type Route struct {
	Uri    string
	Method string
	Function func(w http.ResponseWriter, r *http.Request)
}