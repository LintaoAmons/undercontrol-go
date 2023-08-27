package usecase

import (
	"github.com/LintaoAmons/undercontrol/src/persistence/entgo"
	persistency "github.com/LintaoAmons/undercontrol/src/persistence/entgo/account"

	"github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
)

// HACK:
// Usecase should
// 1. Either a transparent call of one domain service method or domain method
// 2. Or an orchestration of several domains
// Transactional at this level should be catered very carefully since it may involves multi schemas or multi datasources

// 关于跨层调用: 主要目的还是让数据变动可控，增删改的操作能集中在某个层一览无余，而不用担心数据在某个未知角落被奇怪的操作修改
// 如果是 读: 外部知道核心的API，直接依赖核心也无所谓.而且也没有Transction需要操心
//
//	比如 一个 controler, 作为系统最外层，如果只是 list 一些资源，直接依赖 Repo 调用方法也行
//
// 如果涉及写： 就要严格的一层一层的调用, 要重点考虑如何保证数据一致性，如何实现Transaction
// 这时候，读写分离就能很明显地强制要求每个开发对读写进行分类，并按照各自的规则进行依赖以及调用
type AccountUsecase struct {
	repo    account.AccountRepository
	service account.AccountService
}

// TODO: refactor the builder method. Maybe wired?
func NewAccountUsecase() *AccountUsecase {
	client := entgo.InitDBClient()
	repo := persistency.NewAccountEntRepo(client)
	histRepo := persistency.NewAccountHistoryEntRepo(client)
	// TODO: Account Factory, how to init only one instance of each type
	service := account.NewAccountService(repo, histRepo, common.DefaultTxManager{})
	return &AccountUsecase{
		repo:    repo,
		service: service,
	}
}

func (au *AccountUsecase) FindAll() []*account.Account {
	return au.repo.FindAll()
}

type CreateAccountCommand = account.CreateAccountCommand

func (au *AccountUsecase) Add(c CreateAccountCommand) {
	au.service.CreateNewAccount(c)
}

func (au *AccountUsecase) Get(name string) (*account.Account, error) {
	result, err := au.repo.Get(name)

	return result, err
}

type DepositCommand = account.DepositCommand

func (au *AccountUsecase) Deposit(dc DepositCommand) {
	au.service.Deposit(dc)
}

type WithdarwCommand = account.WithdarwCommand

func (au *AccountUsecase) Withdarw(command WithdarwCommand) {
	au.service.Withdarw(command)
}

type Calibratecommand = account.CalibrateCommand

func (au *AccountUsecase) Calibrate(command Calibratecommand) {
	au.service.Calibrate(command)
}
