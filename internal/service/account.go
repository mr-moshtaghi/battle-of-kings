package service

import (
	"battle-of-kings/internal/entity"
	"battle-of-kings/internal/repository"
	"context"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(rep repository.AccountRepository) *AccountService {
	return &AccountService{accounts: rep}
}

func (a *AccountService) UpdateOrCreate(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
