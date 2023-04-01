package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"kv-mem/internal/server"
	"kv-mem/internal/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestServer_GetValue(t *testing.T) {
	db := store.New(10, 10)
	if err := db.Set("key", "value"); err != nil {
		require.Nil(t, err)
	}

	srv := server.New("8080", db)
	ts := httptest.NewServer(srv.Handler)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL+"/store/key", nil)
	require.Nil(t, err)

	res, err := hc.Do(req)
	require.Nil(t, err)

	body, err := io.ReadAll(res.Body)
	require.Nil(t, err)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "value", string(body))
}

func TestServer_SetValue(t *testing.T) {
	db := store.New(10, 10)
	srv := server.New("8080", db)
	ts := httptest.NewServer(srv.Handler)
	defer ts.Close()

	val := "value"
	req, err := http.NewRequest(http.MethodPut, ts.URL+"/store/key", strings.NewReader(val))
	require.Nil(t, err)

	res, err := hc.Do(req)
	require.Nil(t, err)

	body, err := io.ReadAll(res.Body)
	require.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, res.StatusCode)
	assert.Equal(t, "", string(body))
}

var hc = http.Client{Timeout: 2 * time.Second}
