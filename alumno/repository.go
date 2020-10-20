//Aqui se hacen las consultas y las transacciones
package alumno

import "database/sql"

//Tendremos un metodos en la interface para implementar en una estructura
type Repository interface {
	GetAlumnoById(studentId int) (*Student, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetAlumnoById(studentId int) (*Student, error) {
	const sql = `select * from student where STUDENT_ID=?`
	row := repo.db.QueryRow(sql, studentId)
	//Mapaear todos los valores de la consulta sql en el puntero +Student
	student := &Student{}
	//La funcion scan ayuda a mapear los resultados a una entidad determinada en la entidad student
	err := row.Scan(&student.Student_id, &student.Student_name,
		&student.Student_country, &student.Student_gender,
		&student.Student_email, &student.Student_password,
		&student.Student_semester) //Poniendo valores a la estructura segun la consulta
	if err != nil {
		panic(err)
	}
	return student, err
}
