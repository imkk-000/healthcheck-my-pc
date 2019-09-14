package api

import (
	"healthcheck/server/service"
	"io/ioutil"
	"net/http"
)

type HealthcheckAPI struct {
	HealthcheckServiceImpl service.HealthcheckService
}

func (api HealthcheckAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validate uri
	switch r.URL.Path {
	case "/healthcheck":
	case "/healthcheck/":
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// validate method
	switch r.Method {
	case http.MethodPost:
		// get body
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(body) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		api.HealthcheckServiceImpl.SaveDataToCache(body)
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		w.Write(api.HealthcheckServiceImpl.GetCacheData())
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
