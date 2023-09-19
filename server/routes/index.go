package routes

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func Index(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	body, err := os.ReadFile("../client/index.html")
	if err != nil {
		panic(err)
	}

	response.WriteHeader(http.StatusOK)
	response.Write(body)
}
