package store_test

import (
	"github.com/stretchr/testify/assert"
	"kv-mem/internal/store"
	"sync"
	"testing"
)

func TestSet_max_key_len_exceeded(t *testing.T) {
	s := store.New(5, 10)
	_, err := s.Set("123456", "bar", "")
	assert.ErrorIs(t, err, store.ErrMaxKeyLen)
}

func TestSet_max_val_len_exceeded(t *testing.T) {
	s := store.New(5, 9)
	_, err := s.Set("foo", "0123456789", "")
	assert.ErrorIs(t, err, store.ErrMaxValLen)
}

func TestSet_etag_mismatch(t *testing.T) {
	s := store.New(10, 10)
	_, _ = s.Set("foo", "bar", "")
	_, err := s.Set("foo", "bar", "incorrect")

	assert.ErrorIs(t, err, store.ErrETagMismatch)
}

func TestSetAndGet_success(t *testing.T) {
	s := store.New(10, 10)

	etag, _ := s.Set("foo", "bar", "")
	assert.Equal(t, store.Hash("bar"), etag)

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "bar", got)
}

func TestSetAndGet_overwrite_value(t *testing.T) {
	s := store.New(10, 10)

	etag, _ := s.Set("foo", "bar", "")
	_, err := s.Set("foo", "buzz", etag)
	assert.Nil(t, err)

	got, err := s.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, "buzz", got)
}

func TestGet_nonexistent_key(t *testing.T) {
	s := store.New(10, 10)
	_, err := s.Get("foo")
	assert.ErrorIs(t, err, store.ErrNotFound)
}

func TestSet_concurrency(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()
	s := store.New(10, 10)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.Set("foo", "bar", "")
			assert.Nil(t, err)
		}()
	}
}

func TestGet_concurrency(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()
	s := store.New(10, 10)
	_, _ = s.Set("foo", "bar", "")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := s.Get("foo")
			assert.Nil(t, err)
		}()
	}
}
