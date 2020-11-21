package routers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/RestApiELearningLP2/database"
	"github.com/RestApiELearningLP2/models"
	"github.com/RestApiELearningLP2/usuario"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Valor usado de Email usado en todos los EndPoints*/
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints*/
var IDUsuario string

/*ProcesosToken Proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("XDXDXD_token_XDXDXD")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	databaseConnection := database.InitDB()
	defer databaseConnection.Close()
	var repository = usuario.NewRepository(databaseConnection)

	if err == nil {
		var encontradoBool bool = false
		claims1, _ := tkn.Claims.(jwt.MapClaims)
		if tkn.Valid {
			fmt.Println(claims1)
			IDUsuario1 := claims1["_id"]
			Email1 := claims1["email"]
			fmt.Println(IDUsuario1)
			fmt.Println(Email1)
		}
		_, encontrado, errr := repository.ChequeoExisteUsuario(&claims.Email)
		if errr != nil {
			return claims, false, string(""), errr
		}
		if encontrado == 1 {
			Email = claims.Email
			IDUsuario = claims.ID
			encontradoBool = true
		}
		return claims, encontradoBool, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
