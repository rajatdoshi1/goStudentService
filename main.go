package main

import (
	"Practise/goStudentService/StudentAccount"
	"Practise/goStudentService/StudentModel"
	"Practise/goStudentService/webservice"
	"context"
	"flag"
	"fmt"
	"net/http"
)

func main() {
	repo := StudentAccount.NewRepository()

	s1 := StudentModel.Student{Name: "S1", Class: 1, RollNo: 1}
	s2 := StudentModel.Student{Name: "S2", Class: 2, RollNo: 2}
	s3 := StudentModel.Student{Name: "S3", Class: 3, RollNo: 3}
	s4 := StudentModel.Student{Name: "S1", Class: 4, RollNo: 4}
	s5 := StudentModel.Student{Name: "S5", Class: 5, RollNo: 5}
	s6 := StudentModel.Student{Name: "S6", Class: 6, RollNo: 6}

	repo.AddNewStudent(s1)
	repo.AddNewStudent(s2)
	repo.AddNewStudent(s3)
	repo.AddNewStudent(s4)
	repo.AddNewStudent(s5)
	repo.AddNewStudent(s6)

	student, _ := repo.GetStudentById(4)
	fmt.Printf("%s", student.Name)

	var httpAddr = flag.String("http", ":8080", "HTTP server")
	
	flag.Parse()
	backgroundContext := context.Background()

	errs := make(chan error)

	srv := StudentAccount.NewHandler(repo)
	endpoints := webservice.MakeEndpoints(srv)

	fmt.Println("listening on port", *httpAddr)
	handler := webservice.NewHTTPServer(backgroundContext, endpoints)
	errs <- http.ListenAndServe(*httpAddr, handler)
}
