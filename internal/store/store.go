package store

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrorMaxKeyLen = errors.New("max key length exceeded")
	ErrorMaxValLen = errors.New("max val length exceeded")
)

func New(maxKeyLen, maxValLen int) *Store {
	return &Store{
		maxKeyLen: maxKeyLen,
		maxValLen: maxValLen,
		data:      make(map[string]string),
	}
}

// Store is a simple key-value store.
type Store struct {
	maxKeyLen int
	maxValLen int
	data      map[string]string
}

// Set stores the given value under the given key.
func (s *Store) Set(key, value string) error {
	if len(key) > s.maxKeyLen {
		return fmt.Errorf("%w: key: %s", ErrorMaxKeyLen, key)
	}

	if len(value) > s.maxValLen {
		return fmt.Errorf("%w: value: %s", ErrorMaxValLen, value)
	}

	s.data[key] = value
	return nil
}

// Get returns the value stored under the given key.
// If the key does not exist, it returns an error.
func (s *Store) Get(key string) (string, error) {
	if _, ok := s.data[key]; !ok {
		return "", fmt.Errorf("%w: key: %s", ErrNotFound, key)
	}
	return s.data[key], nil
}
