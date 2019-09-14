package service

type HealthcheckService interface {
	SaveDataToCache(data []byte)
	GetCacheData() []byte
}

type HealthcheckServiceImpl struct {
	CacheData []byte
}

func (service *HealthcheckServiceImpl) SaveDataToCache(data []byte) {
	service.CacheData = data
}

func (service HealthcheckServiceImpl) GetCacheData() []byte {
	return service.CacheData
}
