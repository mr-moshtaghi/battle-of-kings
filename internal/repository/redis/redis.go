package redis

import (
	"github.com/redis/rueidis"
)

func NewRedisClient(addr string) (rueidis.Client, error) {
	return rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{addr},
	})
}
