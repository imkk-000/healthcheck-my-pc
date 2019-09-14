package api_test

import (
	"bytes"
	. "healthcheck/server/api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const URL_TEST = "http://test.local"

func TestHealthcheckHandlerPostWithEmptyDataShouldBeStatusBadRequest(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, URL_TEST+"/healthcheck", bytes.NewBuffer([]byte("")))
	w := httptest.NewRecorder()

	healthcheckAPI := &HealthcheckAPI{
		HealthcheckServiceImpl: &mockHealthcheckService{},
	}
	healthcheckAPI.HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHealthcheckHandlerPostWithEmptyURLShouldBeStatusBadRequest(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()

	healthcheckAPI := &HealthcheckAPI{
		HealthcheckServiceImpl: &mockHealthcheckService{},
	}
	healthcheckAPI.HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHealthcheckHandlerGetDataWithSlashURLShouldBeRespData(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, URL_TEST+"/healthcheck/", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()

	healthcheckAPI := &HealthcheckAPI{
		HealthcheckServiceImpl: &mockHealthcheckService{},
	}
	healthcheckAPI.HealthcheckHandler(w, r)
	body, err := ioutil.ReadAll(w.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, []byte("{}"), body)
}

func TestHealthcheckHandlerPostDataShouldBeStatusOK(t *testing.T) {
	r, _ := http.NewRequest(http.MethodPost, URL_TEST+"/healthcheck", bytes.NewBuffer([]byte("{}")))
	w := httptest.NewRecorder()

	healthcheckAPI := &HealthcheckAPI{
		HealthcheckServiceImpl: &mockHealthcheckService{},
	}
	healthcheckAPI.HealthcheckHandler(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
