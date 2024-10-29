package repository

import (
	"battle-of-kings/internal/entity"
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("entry not found")
)

type AccountRepository interface {
	CommonBehaviour[entity.Account]
}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, ent entity.Entity) error
}
