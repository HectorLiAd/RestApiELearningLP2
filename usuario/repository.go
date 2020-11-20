//Aqui se hacen las consultas y las transacciones
package usuario

import "database/sql"

//Tendremos un metodos en la interface para implementar en una estructura
type Repository interface {
	ChequeoExisteUsuario(email *string) (*Usuario, int, error)
	InsertoRegistro(params *registerUserRequest) (string, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) ChequeoExisteUsuario(email *string) (*Usuario, int, error) {
	contCorreo := -1
	usuario := &Usuario{}
	const queryStr = `SELECT * from USUARIO WHERE EMAIL = ?`
	row := repo.db.QueryRow(queryStr, email)
	row.Scan(&usuario.Usuario_id, &usuario.Usuario_nombre,
		&usuario.Usuario_email, &usuario.Usuario_password,
		&usuario.Usuario_avatar)

	const queryStrCont = `SELECT COUNT(EMAIL) FROM USUARIO WHERE EMAIL = ?`
	rowCont := repo.db.QueryRow(queryStr, email)
	rowCont.Scan(&contCorreo)
	return usuario, contCorreo, nil
}

func (repo *repository) InsertoRegistro(params *registerUserRequest) (string, error) {
	var dato string = "success"
	const queryStr = `INSERT INTO USUARIO VALUES(?, ?, ?, ?, ?)`
	_, err := repo.db.Exec(
		queryStr, params.Usuario_id,
		params.Usuario_nombre, params.Usuario_email,
		params.Usuario_password, params.Usuario_avatar)

	return dato, err
}
