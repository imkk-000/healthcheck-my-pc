package util_test

import (
	. "healthcheck/client/util"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendDataShoubeBeStatusBadRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()
	err := SendData(ts.URL, "application/json", []byte(""))

	assert.NotNil(t, err)
}

func TestSendDataShoubeBeStatusOK(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()
	err := SendData(ts.URL, "application/json", []byte(""))

	assert.Nil(t, err)
}
