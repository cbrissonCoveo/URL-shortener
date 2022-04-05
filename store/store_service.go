package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Define struct wrapper arround raw Redis Client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 24 * time.Hour

// Fetch original URL from shortened URL
func RetrieveInitialUrl(shortUrl string) string {
	result, err  := storeService.redisClient.Get(shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s", err, shortUrl))
	}
	return result
}

// Save the mapping between originalUrl and shortUrl
func SaveUrlMapping(shortUrl, originalUrl, userId string) {
	err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error %v - shortUrl %s - originalUrl %s", err, shortUrl, originalUrl))
	}
}



func InitializeStore() *StorageService{
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB : 0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error while initiating Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService

}