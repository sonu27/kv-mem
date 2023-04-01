package store_test

import (
	"github.com/stretchr/testify/assert"
	"kv-mem/internal/store"
	"testing"
)

func TestSet_max_key_len_exceeded(t *testing.T) {
	s := store.New(5, 10)
	err := s.Set("123456", "bar")
	assert.ErrorIs(t, err, store.ErrorMaxKeyLen)
}

func TestSet_max_val_len_exceeded(t *testing.T) {
	s := store.New(5, 9)
	err := s.Set("foo", "0123456789")
	assert.ErrorIs(t, err, store.ErrorMaxValLen)
}
func TestSetAndGet_success(t *testing.T) {
	s := store.New(10, 10)
	_ = s.Set("foo", "bar")

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", got)
}

func TestSetAndGet_overwrite_value(t *testing.T) {
	s := store.New(10, 10)
	_ = s.Set("foo", "bar")
	_ = s.Set("foo", "buzz")

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "buzz", got)
}

func TestGet_nonexistent_key(t *testing.T) {
	s := store.New(10, 10)
	_, err := s.Get("foo")
	assert.ErrorIs(t, err, store.ErrNotFound)
}
