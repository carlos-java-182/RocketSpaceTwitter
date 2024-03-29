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
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Registro)).Methods(("POST"))
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods(("GET"))
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods(("PUT"))
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods(("POST"))
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods(("GET"))
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods(("DELETE"))
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(routers.SubirAvatar)).Methods(("POST"))
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods(("GET"))
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(routers.SubirBanner)).Methods(("POST"))
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods(("GET"))
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(routers.AltaRelacion)).Methods(("POST"))
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(routers.BajaRelacion)).Methods(("DELETE"))
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(routers.ConsultaRelacion)).Methods(("GET"))
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(routers.ListaUsuarios)).Methods(("GET"))
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(routers.LeoTweetsSeguidores)).Methods(("GET"))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
