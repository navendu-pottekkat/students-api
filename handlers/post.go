package handlers

import (
	"net/http"

	"github.com/navendu-pottekkat/students-api/data"
)

// AddStudent adds a new student
func (s *Students) AddStudent(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST Student")
	stud := r.Context().Value(KeyStudent{}).(data.Student)
	data.AddStudent(&stud)
}
