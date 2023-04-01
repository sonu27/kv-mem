package store_test

import (
	"github.com/stretchr/testify/assert"
	"kv-mem/internal/store"
	"testing"
)

func TestSetAndGet_success(t *testing.T) {
	s := store.New()
	s.Set("foo", "bar")

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", got)
}

func TestSetAndGet_overwrite_value(t *testing.T) {
	s := store.New()
	s.Set("foo", "bar")
	s.Set("foo", "buzz")

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "buzz", got)
}

func TestGet_nonexistent_key(t *testing.T) {
	s := store.New()
	_, err := s.Get("foo")
	assert.ErrorIs(t, err, store.ErrNotFound)
}
