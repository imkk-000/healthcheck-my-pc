package api_test

import (
	"bytes"
	. "healthcheck/server/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckHandlerPostPushInfoWithEmptyData(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "", bytes.NewBuffer([]byte("")))
	w := httptest.NewRecorder()

	HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHealthcheckHandlerGetPushInfo(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()

	HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHealthcheckHandlerPostPushInfo(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()

	HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
