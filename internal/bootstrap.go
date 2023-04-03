package internal

import (
	"kv-mem/internal/server"
	"kv-mem/internal/store"
	"log"
)

func Bootstrap(port string) error {
	kvStore := store.New(100, 100)

	srv := server.New(port, kvStore, 10000)

	log.Printf("server started on http://localhost:%s", port)
	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
