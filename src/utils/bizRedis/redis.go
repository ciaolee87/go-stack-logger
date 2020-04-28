package bizRedis

import "github.com/go-redis/redis"

func NewClient(server string, pw string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               server,
		Dialer:             nil,
		OnConnect:          nil,
		Password:           pw,
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	return client
}
