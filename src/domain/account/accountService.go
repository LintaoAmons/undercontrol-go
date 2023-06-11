package account

import (
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/Rhymond/go-money"
	"github.com/pterm/pterm"
)

type AccountService struct {
	repo   AccountRepository
	logger *pterm.Logger
}

func NewAccountService(repo AccountRepository) *AccountService {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithMaxWidth(200)
	return &AccountService{
		repo:   repo,
		logger: logger,
	}
}

type CreateAccountCommand struct {
	Name     string
	Amount   float64
	Currency string
}

func (cac CreateAccountCommand) create() *Account {
	return &Account{
		Name:   cac.Name,
		Amount: money.NewFromFloat(cac.Amount, cac.Currency),
		Audit:  common.DefaultAudit(),
	}
}

func (as *AccountService) CreateNewAccount(c *CreateAccountCommand) {
	as.repo.Save(c.create())
}
