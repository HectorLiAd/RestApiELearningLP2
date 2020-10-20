package alumno

type Student struct {
	Student_id       int    `json:"student_id"`
	Student_name     string `json:"student_name"`
	Student_country  string `json:"student_country"`
	Student_gender   string `json:"student_gender"`
	Student_email    string `json:"student_email"`
	Student_password string `json:"student_password"`
	Student_semester string `json:"student_semester"`
}
