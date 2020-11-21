package usuario

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t *Usuario) (string, error) {
	miClave := []byte("XDXDXD_token_XDXDXD") //Creando clave privada
	claims := jwt.MapClaims{}
	claims["email"] = t.Usuario_email
	claims["nombre"] = t.Usuario_nombre
	claims["_id"] = t.Usuario_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// payload := jwt.MapClaims{
	// 	"email":  t.Usuario_email,
	// 	"nombre": t.Usuario_nombre,
	// 	"_id":    t.Usuario_id,
	// 	"exp":    time.Now().Add(time.Hour * 24).Unix(),
	// }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims) //El goritmo para encriptar "header"
	tokenStr, err := token.SignedString(miClave)                //Firmando con el string de firma
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
