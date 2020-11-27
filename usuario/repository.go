package usuario

import (
	"database/sql"
	"fmt"
)

//Repository Tendremos un metodos en la interface para implementar en una estructura
type Repository interface {
	ChequeoExisteUsuario(email *string) (*Usuario, int, error)
	InsertoRegistro(params *registerUserRequest) (string, error)
}

type repository struct {
	db *sql.DB
}

/*NewRepository creara el repositorio*/
func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) ChequeoExisteUsuario(email *string) (*Usuario, int, error) {
	contCorreo := 2
	usuario := &Usuario{}
	const queryStr = `SELECT * from USUARIO WHERE EMAIL = ?`
	row := repo.db.QueryRow(queryStr, email)
	row.Scan(&usuario.UsuarioID, &usuario.UsuarioNombre,
		&usuario.UsuarioEmail, &usuario.UsuarioPassword,
		&usuario.UsuarioAvatar)

	const queryStrCont = `SELECT COUNT(EMAIL) FROM USUARIO WHERE EMAIL = ?`
	rowCont := repo.db.QueryRow(queryStrCont, email)
	rowCont.Scan(&contCorreo)
	fmt.Println(contCorreo)
	return usuario, contCorreo, nil
}

func (repo *repository) InsertoRegistro(params *registerUserRequest) (string, error) {
	var dato string = "success"
	const queryStr = `INSERT INTO USUARIO VALUES(?, ?, ?, ?, ?)`
	_, err := repo.db.Exec(
		queryStr, params.UsuarioID,
		params.UsuarioNombre, params.UsuarioEmail,
		params.UsuarioPassword, params.UsuarioAvatar)

	return dato, err
}
