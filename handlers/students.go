package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/navendu-pottekkat/students-api/data"
)

// Details is a http handler
type Students struct {
	l *log.Logger
}

type KeyStudent struct{}

func NewStudents(l *log.Logger) *Students {
	return &Students{l}
}

// GetStudents returns the students
func (s *Students) GetStudents(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle GET Students")

	ls := data.GetStudents()

	// Serialize to JSON
	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddStudent adds a new student
func (s *Students) AddStudent(rw http.ResponseWriter, r *http.Request) {
	s.l.Println("Handle POST Student")
	stud := r.Context().Value(KeyStudent{}).(data.Student)
	data.AddStudent(&stud)
}

func (s Students) MiddlewareValidateStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		stud := data.Student{}

		err := stud.FromJSON(r.Body)

		if err != nil {
			s.l.Println("[Error] deserialising student", err)
			http.Error(
				rw,
				"Error reading student",
				http.StatusBadRequest,
			)
			return
		}

		// Validate the student
		err = stud.Validate()
		if err != nil {
			s.l.Println("[ERROR] validating student", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating student: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// Add the student to the context
		ctx := context.WithValue(r.Context(), KeyStudent{}, stud)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(rw, r)
	})
}
