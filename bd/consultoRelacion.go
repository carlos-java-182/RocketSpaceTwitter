package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/carlos-java-182/RocketSpaceTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre 2 usuario*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("RocketSpaceTwitter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuario":         t.UsuarioID,
		"usuariorelacion": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	fmt.Print(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}

	return true, nil

}
