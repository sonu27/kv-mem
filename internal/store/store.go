package store

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrNotFound     = errors.New("key not found")
	ErrMaxKeyLen    = errors.New("max key length exceeded")
	ErrMaxValLen    = errors.New("max val length exceeded")
	ErrETagMismatch = errors.New("etag does not match")
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
	mu        sync.Mutex
}

// Set stores the given value under the given key.
func (s *Store) Set(key, value, etag string) (string, error) {
	if len(key) > s.maxKeyLen {
		return "", fmt.Errorf("%w: key: %s", ErrMaxKeyLen, key)
	}

	if len(value) > s.maxValLen {
		return "", fmt.Errorf("%w: value: %s", ErrMaxValLen, value)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if etag != "" {
		v, ok := s.data[key]
		if !ok {
			return "", fmt.Errorf("%w: key does not exist: %s", ErrETagMismatch, key)
		}

		et := Hash(v)
		if etag != et {
			return "", fmt.Errorf("%w: key: %s", ErrETagMismatch, key)
		}
	}

	s.data[key] = value
	return Hash(value), nil
}

// Get returns the value stored under the given key.
// If the key does not exist, it returns an error.
func (s *Store) Get(key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[key]; !ok {
		return "", fmt.Errorf("%w: %s", ErrNotFound, key)
	}
	return s.data[key], nil
}

func Hash(value string) string {
	etag := md5.Sum([]byte(value))
	return hex.EncodeToString(etag[:])
}
