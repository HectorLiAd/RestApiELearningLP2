//Pondremos la logica del negocio y consumiremos los metodos que estan en repository,go
package usuario

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/RestApiELearningLP2/helper"
)

type Service interface {
	registrarUsuario(params *registerUserRequest) (string, error)
	IntentoLogin(params *loginUserRequest) (*Usuario, error)
	LoginUsuario(params *loginUserRequest) (interface{}, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

/*Registrar Usuario*/
func (s *service) registrarUsuario(params *registerUserRequest) (string, error) {
	//Logica de negocio

	if len(params.Usuario_email) == 0 {
		return "", errors.New("El email es requerido")
	}
	if len(params.Usuario_password) < 6 {
		return "", errors.New("Debe especificar una contraseña de almenos 6 caracteres")
	}
	// Verificando si hay doplicidad en gmail
	_, usuarioCorreo, _ := s.repo.ChequeoExisteUsuario(&params.Usuario_email)
	if usuarioCorreo > 0 || usuarioCorreo == -1 {
		return "", errors.New("Ya existe un usuario registrado con ese email")
	}
	// Encriptando la contraseña
	pwdEncriptado, err := helper.EncriptarPassword(params.Usuario_password)
	if err != nil {
		return "", err
	}
	params.Usuario_password = pwdEncriptado
	// Insertando persona a la BD
	id_usuario, err := s.repo.InsertoRegistro(params)

	if err != nil {
		return "", err
	}

	return id_usuario, err
}

/*Intento Login*/
func (s *service) IntentoLogin(params *loginUserRequest) (*Usuario, error) {
	usuario, encontrado, err := s.repo.ChequeoExisteUsuario(&params.Usuario_email)
	if err != nil {
		return nil, err
	}
	if encontrado <= 0 {
		return nil, errors.New("Usuario no encontrado")
	}
	passwordBytes := []byte(params.Usuario_password)
	passwordBD := []byte(usuario.Usuario_password)
	//Verificando la PWD
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return nil, errors.New("Usuario y/o Contraseña invalidos " + err.Error())
	}
	return usuario, nil
}

func (s *service) LoginUsuario(params *loginUserRequest) (interface{}, error) {
	if len(params.Usuario_email) == 0 {
		return nil, errors.New("El email del usuario es requerido")
	}
	usuario, err := s.IntentoLogin(params)
	if err != nil {
		return usuario, err
	}

	//JWT
	jwtkey, er := jwt.GeneroJWT(usuario)
	if err != nil {
		return jwtkey, errors.New("El email del usuario es requerido" + er.errors())
	}
	resp := RespuestaLogin{
		Token: jwtkey,
	}

	// CUARGAR UNA COOKISSS DEL USUARIO PARA ACCEDER DESDE EL FRONT
	// expirationTime := time.Now().Add(24 * time.Hour) //
	// http.SetCookie(w, &http.Cookie{
	// 	Name: "token",
	// })
	return resp, nil
}
