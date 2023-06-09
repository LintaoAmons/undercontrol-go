package account

import (
	"database/sql"

	domain "github.com/LintaoAmons/undercontrol/src/domain/account"
)

type AccountRepositoryImpl struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) domain.AccountRepository {
	return &AccountRepositoryImpl{db: db}
}

func (ar *AccountRepositoryImpl) Save(a *domain.Account) (id string) {
  // TODO:
	return "1"
}

func (ar *AccountRepositoryImpl) Get(id string) (*domain.Account, error) {
  // TODO:
  return nil, nil
}

func (ar *AccountRepositoryImpl) Find(id string) *domain.Account {
  // TODO:
  return nil
}

func (ar *AccountRepositoryImpl) FindAll() *[]domain.Account {
  // TODO:
  return nil
}
