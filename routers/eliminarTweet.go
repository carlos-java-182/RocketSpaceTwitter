package routers

import (
	"net/http"

	"github.com/carlos-java-182/RocketSpaceTwitter/bd"
)

/*EliminarTweet permite borrar un tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet", http.StatusBadRequest)
		return

	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
