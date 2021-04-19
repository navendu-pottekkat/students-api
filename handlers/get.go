package handlers

import (
	"net/http"

	"github.com/navendu-pottekkat/students-api/data"
)

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
