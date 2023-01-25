package routers

import (
	"net/http"

	"github.com/carlos-java-182/RocketSpaceTwitter/bd"
	"github.com/carlos-java-182/RocketSpaceTwitter/models"
)

/*BajaRelacion da de baja la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get(("id"))
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
