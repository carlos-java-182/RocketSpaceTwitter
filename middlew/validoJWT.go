package middlew

import (
	"net/http"

	"github.com/carlos-java-182/RocketSpaceTwitter/routers"
)

/*Valido JWT permite validar el JWT que nos viene en la petición*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {

			http.Error(w, "Error en el token"+err.Error(), 400)
			return

		}

		next.ServeHTTP(w, r)

	}

}
