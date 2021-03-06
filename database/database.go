package database

import (
	"database/sql"
)

/*InitDB permite crear la conexion a la BD*/
func InitDB() *sql.DB {
	connectionString := "root:hector@tcp(localhost:3306)/E_LEARNING"
	databaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error()) // Error Handling = manejo de errores
	}
	return databaseConnection
}
