package webservice

import (
	student "Practise/goStudentService/StudentModel"
	"context"
)

type Service interface {
	GetStudentById(ctx context.Context, id int) (*student.Student, error)
	GetStudentsByName(ctx context.Context, name string) ([]student.Student, error)
}
