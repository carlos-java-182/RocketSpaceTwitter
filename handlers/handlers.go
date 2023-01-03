package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/carlos-java-182/RocketSpaceTwitter/middlew"
	"github.com/carlos-java-182/RocketSpaceTwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo es escuchar al Servidor*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods(("POST"))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
