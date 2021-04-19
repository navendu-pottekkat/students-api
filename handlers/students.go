package handlers

import (
	"log"
)

// Details is a http handler
type Students struct {
	l *log.Logger
}

type KeyStudent struct{}

func NewStudents(l *log.Logger) *Students {
	return &Students{l}
}
