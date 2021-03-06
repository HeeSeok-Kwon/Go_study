package mydict

import "errors"

// alias
// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	// return "", errors.New("Not Found")
	return "", errNotFound
}

// type Money int
