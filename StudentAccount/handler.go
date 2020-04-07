package StudentAccount

import (
	student "Practise/goStudentService/StudentModel"
	webservice "Practise/goStudentService/webservice"
	"context"
)

type handler struct {
	repostory student.StudentRepository
}

func NewHandler(rep student.StudentRepository) webservice.Service {
	return &handler{
		repostory: rep,
	}
}

func (s handler) GetStudentById(ctx context.Context, id int) (*student.Student, error) {

	student, err := s.repostory.GetStudentById(id)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s handler) GetStudentsByName(ctx context.Context, name string) ([]student.Student, error) {

	students, err := s.repostory.GetStudentsByName(name)

	if err != nil {
		return nil, err
	}

	return students, nil
}
