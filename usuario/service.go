package usuario

import (
	"errors"

	"github.com/RestApiELearningLP2/helper"
	"golang.org/x/crypto/bcrypt"
)

/*Service interface para crear las firmas que se usaran en el enpoint*/
type Service interface {
	registrarUsuario(params *registerUserRequest) (string, error)
	IntentoLogin(params *loginUserRequest) (*Usuario, error)
	LoginUsuario(params *loginUserRequest) (interface{}, error)
}

type service struct {
	repo Repository
}

/*NewService permite crear el servicio*/
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

/*Registrar Usuario*/
func (s *service) registrarUsuario(params *registerUserRequest) (string, error) {
	//Logica de negocio

	if len(params.UsuarioEmail) == 0 {
		return "", errors.New("El email es requerido")
	}
	if len(params.UsuarioPassword) < 6 {
		return "", errors.New("Debe especificar una contraseña de almenos 6 caracteres")
	}
	// Verificando si hay doplicidad en gmail
	_, usuarioCorreo, _ := s.repo.ChequeoExisteUsuario(&params.UsuarioEmail)
	if usuarioCorreo > 0 {
		return "", errors.New("Ya existe un usuario registrado con ese email")
	}
	// Encriptando la contraseña
	pwdEncriptado, err := helper.EncriptarPassword(params.UsuarioPassword)
	if err != nil {
		return "", err
	}
	params.UsuarioPassword = pwdEncriptado
	// Insertando persona a la BD
	idUsuario, err := s.repo.InsertoRegistro(params)

	if err != nil {
		return "", err
	}

	return idUsuario, err
}

/*Intento Login*/
func (s *service) IntentoLogin(params *loginUserRequest) (*Usuario, error) {
	usuario, encontrado, err := s.repo.ChequeoExisteUsuario(&params.UsuarioEmail)
	if err != nil {
		return nil, err
	}
	if encontrado == 0 {
		return nil, errors.New("Usuario no encontrado")
	}
	passwordBytes := []byte(params.UsuarioPassword)
	passwordBD := []byte(usuario.UsuarioPassword)
	//Verificando la PWD
	err = bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return nil, errors.New("Usuario y/o Contraseña invalidos " + err.Error())
	}
	return usuario, nil
}

func (s *service) LoginUsuario(params *loginUserRequest) (interface{}, error) {
	if len(params.UsuarioEmail) == 0 {
		return nil, errors.New("El email del usuario es requerido")
	}
	usuario, err := s.IntentoLogin(params)
	if err != nil {
		return usuario, err
	}

	//JWT
	jwtkey, er := GeneroJWT(usuario)
	if er != nil {
		return "", errors.New("El email del usuario es requerido" + er.Error())
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
