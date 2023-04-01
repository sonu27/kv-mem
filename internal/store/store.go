package store

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Store is a simple key-value store.
type Store struct {
	data map[string]string
}

// Set stores the given value under the given key.
func (s *Store) Set(key, value string) {
	s.data[key] = value
}

// Get returns the value stored under the given key.
// If the key does not exist, it returns an error.
func (s *Store) Get(key string) (string, error) {
	if _, ok := s.data[key]; !ok {
		return "", fmt.Errorf("%w: key: %s", ErrNotFound, key)
	}
	return s.data[key], nil
}
