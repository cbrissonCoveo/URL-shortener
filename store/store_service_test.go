package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortUrl, initialLink, userUUID)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retrievedUrl)
}

func TestStorageInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func init() {
	testStoreService = InitializeStore()
}