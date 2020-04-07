package StudentModel

type Student struct {
	ID     int
	Name   string
	Class  int
	RollNo int
}

type StudentRepository interface {
	AddNewStudent(student Student) error
	GetStudentById(id int) (Student, error)
	GetStudentsByName(name string) ([]Student, error)
}
