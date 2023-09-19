package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func Stammdaten(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	people := []Person{{
		FirstName: "John",
		LastName:  "Doe",
	}, {
		FirstName: "Jane",
		LastName:  "Doe",
	}}

	bytes, err := json.Marshal(people)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(bytes)
}
