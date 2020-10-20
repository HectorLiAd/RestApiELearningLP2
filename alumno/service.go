//Pondremos la logica del negocio y consumiremos los metodos que estan en repository,go
package alumno

type Service interface {
	GetStudentById(param *getStudentByIDRequest) (*Student, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetStudentById(param *getStudentByIDRequest) (*Student, error) {
	//Logica de negocio
	return s.repo.GetAlumnoById(param.StudentID)
}
