package StudentAccount

import (
	student "Practise/goStudentService/StudentModel"
	"errors"
	"fmt"
)

var students = make(map[int]student.Student)
var lastId int = 1

type Repository struct {
}

func NewRepository() student.StudentRepository {
	return &Repository{}
}

func (repo *Repository) AddNewStudent(student student.Student) error {
	student.ID = lastId
	students[lastId] = student
	lastId++
	return nil
}

func (repo *Repository) GetStudentById(id int) (student.Student, error) {
	student := students[id]
	// if student == nil {
	// 	return nil, errors.New(fmt.Sprintf("No Student with Id %d found", id))
	// }
	return student, nil
}

func (repo *Repository) GetStudentsByName(name string) ([]student.Student, error) {
	var searchedStudents []student.Student
	for _, student := range students {
		if student.Name == name {
			searchedStudents = append(searchedStudents, student)
		}
	}
	if len(searchedStudents) == 0 {
		return nil, errors.New(fmt.Sprintf("No students found with name: %s", name))
	}

	return searchedStudents, nil
}
