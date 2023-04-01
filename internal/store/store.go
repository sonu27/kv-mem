package store

func New() *Store {
	return &Store{}
}

// Store is a simple key-value store.
type Store struct {
}

// Set stores the given value under the given key.
func (s *Store) Set(key, value string) {

}

// Get returns the value stored under the given key.
// If the key does not exist, it returns an error.
func (s *Store) Get(key string) (string, error) {
	return "", nil
}
