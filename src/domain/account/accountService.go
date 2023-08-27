package account

import (
	"time"

	"github.com/LintaoAmons/go-money"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	"github.com/pterm/pterm"
	"github.com/shopspring/decimal"
)

// HACK:
// Contain one domain specific BZ logic
// Transaction here most likely to be a single schema transaction
type AccountServiceImpl struct {
	repo      AccountRepository
	histRepo  AccountHistoryRepo
	txManager common.TxManager
	logger    *pterm.Logger
}

func NewAccountService(
	repo AccountRepository,
	histRepo AccountHistoryRepo,
	txManager common.TxManager,
) AccountService {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithMaxWidth(200)
	return &AccountServiceImpl{
		repo:      repo,
		histRepo:  histRepo,
		txManager: txManager,
		logger:    logger,
	}
}

func (as *AccountServiceImpl) CreateNewAccount(command CreateAccountCommand) {
	as.logger.Trace("Start to create account")
	account := command.create()
	as.saveAccountAndHistory(account)
}

func (as *AccountServiceImpl) Calibrate(command CalibrateCommand) {
	target, _ := as.repo.Get(command.Name)
	calibrated := money.NewFromFloat(command.Amount, command.CurrencyCode)
	target.Amount = calibrated
	target.Audit.UpdatedAt = time.Now()
	as.saveAccountAndHistory(target)
}

func (as *AccountServiceImpl) Withdarw(command WithdarwCommand) {
	target, _ := as.repo.Get(command.Name)
	added, _ := target.Amount.Add(money.NewFromFloat(-command.Amount, target.Amount.Currency().Code))
	target.Amount = added
	target.Audit.UpdatedAt = time.Now()

	as.saveAccountAndHistory(target)
}

func (as *AccountServiceImpl) Deposit(command DepositCommand) {
	target, _ := as.repo.Get(command.Name)
	added, _ := target.Amount.Add(money.NewFromFloat(command.Amount, target.Amount.Currency().Code))
	target.Amount = added
	target.Audit.UpdatedAt = time.Now()
	as.saveAccountAndHistory(target)
}

func (as *AccountServiceImpl) Delete(command DeleteCommand) {
	panic("to be implement") // TODO:
}

func (as *AccountServiceImpl) saveAccountAndHistory(a *Account) {

	tx := as.txManager.Begin()
	as.repo.Save(tx, a)
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	as.histRepo.Save(tx, a.ToHistory(nil))
	tx.Commit()
}

func (cac CreateAccountCommand) create() *Account {
	account := &Account{
		Name:   cac.Name,
		Amount: money.New(decimal.NewFromFloat(cac.Amount), cac.Currency),
		Audit:  common.DefaultAudit(),
	}
	return account
}

