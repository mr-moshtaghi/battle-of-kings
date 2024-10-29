package repository

import (
	"battle-of-kings/internal/entity"
	"github.com/redis/rueidis"
)

type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
