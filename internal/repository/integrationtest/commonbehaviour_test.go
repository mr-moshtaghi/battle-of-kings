package integrationtest

import (
	"battle-of-kings/internal/entity"
	"battle-of-kings/internal/repository"
	"battle-of-kings/internal/repository/redis"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", redisPort))
	assert.NoError(t, err)
	fmt.Println("Redis is now connected")
	ctx := context.Background()
	cb := repository.NewRedisCommonBehaviour[testType](redisClient)
	err = cb.Save(ctx, &testType{
		ID:   "99",
		Name: "khodam",
	})
	assert.NoError(t, err)
	val, err := cb.Get(ctx, entity.NewID("testType", "99"))
	assert.Equal(t, "khodam", val.Name)
	assert.Equal(t, "99", val.ID)

	err = cb.Save(ctx, &testType{
		ID:   "99",
		Name: "Sajjad_khodam",
	})
	assert.NoError(t, err)
	val, err = cb.Get(ctx, entity.NewID("testType", "99"))
	assert.Equal(t, "Sajjad_khodam", val.Name)
	assert.Equal(t, "99", val.ID)

	val, err = cb.Get(ctx, entity.NewID("testType", "100"))
	assert.ErrorIs(t, repository.ErrNotFound, err)
	redisClient.Close()
}
