package main

import (
	"bytes"
	"github.com/muerewa/clipbin/internal/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCommonHeader(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	commonHeaders(next).ServeHTTP(rr, r)
	rs := rr.Result()
	expectedValue := "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	assert.Equal(t, rs.Header.Get("Content-Security-Policy"), expectedValue)

	expectedValue = "nosniff"
	assert.Equal(t, rs.Header.Get("X-Content-Type-Options"), expectedValue)

	expectedValue = "0"
	assert.Equal(t, rs.Header.Get("X-XSS-Protection"), expectedValue)

	expectedValue = "Go"
	assert.Equal(t, rs.Header.Get("Server"), expectedValue)

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	assert.Equal(t, string(body), "OK")
}
