package controllers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostRequestStatusOK(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader(`{"description": "test"}`))
	w := httptest.NewRecorder()

	GetJSON(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostRequest(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("POST", "/test", strings.NewReader(`{"description": "test"}`))
	r.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()

	GetJSON(w, r)

	assert.Equal(t, []byte(`{"description": "test"}`), w.Body.Bytes())
}
