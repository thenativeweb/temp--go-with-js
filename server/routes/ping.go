package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Ping(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("pong"))
}
