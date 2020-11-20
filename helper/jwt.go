package helper

import (
	"time"

	"github.com/RestApiELearningLP2/usuario"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t usuario.Usuario) (string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodefacebook") //Creando clave privada
	payload := jwt.MapClaims{
		"email":  t.Usuario_email,
		"nombre": t.Usuario_nombre,
		"_id":    t.Usuario_id,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //El goritmo para encriptar "header"
	tokenStr, err := token.SignedString(miClave)                //Firmando con el string de firma
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
