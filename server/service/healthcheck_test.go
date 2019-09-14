package service_test

import (
	. "healthcheck/server/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheckServiceGetCacheData(t *testing.T) {
	healthcheckService := &HealthcheckServiceImpl{
		CacheData: []byte("{}"),
	}

	cacheData := healthcheckService.GetCacheData()

	assert.NotNil(t, cacheData)
	assert.NotEmpty(t, cacheData)
	assert.Equal(t, cacheData, []byte("{}"))
}

func TestHealthcheckServiceSaveDataToCache(t *testing.T) {
	healthcheckService := &HealthcheckServiceImpl{}

	healthcheckService.SaveDataToCache([]byte("{}"))

	assert.NotNil(t, healthcheckService.CacheData)
	assert.NotEmpty(t, healthcheckService.CacheData)
	assert.Equal(t, healthcheckService.CacheData, []byte("{}"))
}
