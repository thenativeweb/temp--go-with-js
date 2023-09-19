package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thenativeweb/stammdaten/routes"
)

func main() {
	router := httprouter.New()

	router.GET("/ping", routes.Ping)
	router.GET("/stammdaten", routes.Stammdaten)
	router.GET("/", routes.Index)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
