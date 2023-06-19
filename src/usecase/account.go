package usecase

import (
	"time"

	"github.com/LintaoAmons/go-money"
	accountP "github.com/LintaoAmons/undercontrol/src/persistence/account"

	"github.com/LintaoAmons/undercontrol/src/domain/account"
)

type AccountUsecase struct {
	repo    account.AccountRepository
	service *account.AccountService
}

func NewAccountUsecase() *AccountUsecase {
	repo := accountP.NewAccountRepository()
	service := account.NewAccountService(repo)
	return &AccountUsecase{
		repo:    repo,
		service: service,
	}
}

func (au *AccountUsecase) FindAll() []*account.Account {
	return au.repo.FindAll()
}

type CreateAccountCommand = account.CreateAccountCommand

func (au *AccountUsecase) Add(c *CreateAccountCommand) {
	au.service.CreateNewAccount(c)
}

func (au *AccountUsecase) Get(name string) (*account.Account, error) {
	result, err := au.repo.Get(name)

	return result, err
}

type DepositCommand struct {
	Name   string
	Amount float64
}

func (au *AccountUsecase) Deposit(dc *DepositCommand) {
	// TODO: save deposit history
	target, _ := au.repo.Get(dc.Name)
	added, _ := target.Amount.Add(money.NewFromFloat(dc.Amount, target.Amount.Currency().Code))
	target.Amount = added
	target.Audit.UpdatedAt = time.Now()
	au.repo.Save(target)
}

type WithdarwCommand = DepositCommand

func (au *AccountUsecase) Withdarw(dc *WithdarwCommand) {
	// TODO: save Withdarw history
	target, _ := au.repo.Get(dc.Name)
	added, _ := target.Amount.Add(money.NewFromFloat(-dc.Amount, target.Amount.Currency().Code))
	target.Amount = added
	target.Audit.UpdatedAt = time.Now()
	au.repo.Save(target)
}
