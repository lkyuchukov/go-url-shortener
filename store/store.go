package store

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	store = &Store{}
	ctx   = context.Background()
)

const CacheDuration = 1 * time.Hour

type Store struct {
	redis *redis.Client
}

func Init() *Store {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("could not initialize Redis: ", err)
	}

	log.Println("Redis initialized!")
	store.redis = redisClient

	return store
}

func SaveShortUrl(shortUrl string, longUrl string) {
	if err := store.redis.Set(ctx, shortUrl, longUrl, CacheDuration).Err(); err != nil {
		log.Panicln("failed to save url %v: %v", longUrl, err)
	}
}

func RetrieveLongUrl(shortUrl string) string {
	result, err := store.redis.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Panicln("failed to retrieve url %v: %v", shortUrl, err)
	}

	return result
}
