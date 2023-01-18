package middlew

import (
	"net/http"

	"github.com/carlos-java-182/RocketSpaceTwitter/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return

		}

		next.ServeHTTP(w, r)
	}

}
