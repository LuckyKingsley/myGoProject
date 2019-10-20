package redistest

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

func Redistestfunc() {
	redisdbClient := getRedisClient()
	//key := "test"
	//value := "test"
	//redisdb.Set(key, value, 0)
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				redisdbClient.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				redisdbClient.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, IdleConns: %d\n", redisdbClient.PoolStats().TotalConns, redisdbClient.PoolStats().IdleConns)
		}()
	}
	wg.Wait()
}

/**

 */
func getRedisClient() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}
