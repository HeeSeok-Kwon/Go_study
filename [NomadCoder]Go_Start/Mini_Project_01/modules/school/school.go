package school

import (
	"errors"
	"fmt"
)

// School struct - private
type School struct {
	name    string
	student int
}

var errOver = errors.New("Exceeded admission quota")

// Establish of School
func Establish(name string) *School {
	school := School{name: name, student: 200}
	return &school
}

// Admission of X number of Freshmans
func (s *School) Admission(freshman int) error {
	if s.student+freshman > 500 {
		return errOver
	}
	s.student += freshman
	return nil
}

// Show X number of students in our school
func (s School) Students() int {
	return s.student
}

// Graduated X number of students
func (s *School) Graduated(graduate int) {
	s.student -= graduate
}

// Rename of school
func (s *School) Rename(newName string) {
	s.name = newName
}

// Name of school
func (s School) Name() string {
	return s.name
}

func (s School) String() string {
	return fmt.Sprint(s.Name(), " school has ", s.Students(), " students")
}
