package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/carlos-java-182/RocketSpaceTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosTodos Lee los usuarios registrados en el sistema, si recibe "R" en quienes trae solo los que se relacionaron con migo*/
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("RocketSpaceTwitter")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()

	findOptions.SetSkip((page-1)*20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)`+search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil{
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil{
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relacion
		r.UsuarioID=ID
		r.UsuarioRelacionID=s.ID.Hex()

		incluir = false

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false{
			incluir = true
		}
		if tipo == "follow" && encontrado == true{
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir==true{
			s.Password=""
			s.Biografia=""
			s.SitioWeb=""
			s.Ubicacion=""
			s.Banner=""
			s.Email=""
		
			results = append(results, &s)

		}
	}

	err = cur.Error()
		if err != nil{
			fmt.Println(err.Error())
			return results, false
		}

		cur.Close(ctx)
		return results, true

}