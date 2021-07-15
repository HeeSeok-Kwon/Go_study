package score

import "errors"

type Score map[string]int

var errNotFound = errors.New("This subject could not be Found")
var errSubjectExists = errors.New("This subject already exists")
var errCantUpdate = errors.New("Can't update : The subject doesn't exists")
var errCantDelete = errors.New("Cant Delete : The subject doesn't exists")

// Search a subject
func (s Score) Search(subject string) (int, error) {
	value, exists := s[subject]
	if exists {
		return value, nil
	}
	return -1, errNotFound
}

// Record subject and it's score
func (s Score) Record(subject string, score int) error {
	_, err := s.Search(subject)
	switch err {
	case errNotFound:
		s[subject] = score
	case nil:
		return err
	}
	return nil
}

// Update subject and it's score
func (s Score) Update(subject string, score int) error {
	_, err := s.Search(subject)
	switch err {
	case errNotFound:
		return errCantUpdate
	case nil:
		s[subject] = score
	}
	return nil
}

// Delete subject and it's score
func (s Score) Delete(subject string) error {
	_, err := s.Search(subject)
	switch err {
	case errNotFound:
		return errCantDelete
	case nil:
		delete(s, subject)
	}
	return nil
}
