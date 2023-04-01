package server

import (
	"errors"
	"github.com/go-chi/chi"
	"io"
	"kv-mem/internal/store"
	"log"
	"net/http"
)

func New(port string, store *store.Store) http.Server {
	s := server{store: store}
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(""))
	})
	r.Route("/store", func(r chi.Router) {
		r.Get("/{key}", s.GetValue)
		r.Put("/{key}", s.PutValue)
	})

	return http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}

type server struct {
	store *store.Store
}

func (s *server) GetValue(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	val, err := s.store.Get(key)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write([]byte(val))
}

func (s *server) PutValue(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)
	val := r.Body

	b, err := io.ReadAll(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.store.Set(key, string(b)); err != nil {
		if errors.Is(err, store.ErrMaxKeyLen) || errors.Is(err, store.ErrMaxValLen) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
