package api_test

type mockHealthcheckService struct{}

func (mockHealthcheckService) SaveDataToCache(data []byte) {}

func (mockHealthcheckService) GetCacheData() []byte {
	return []byte("{}")
}
