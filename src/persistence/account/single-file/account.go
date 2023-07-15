package singlefile

import (
	"github.com/LintaoAmons/undercontrol/src/domain/account"
	"github.com/LintaoAmons/undercontrol/src/domain/common"
	singlefile "github.com/LintaoAmons/undercontrol/src/persistence/common/single-file"
)

const ACCOUNT singlefile.Table = "account"

type SingleFileAccountRepo struct {
	db singlefile.SingleFileDB
}

func NewSingleFileAccountRepo(db singlefile.SingleFileDB) account.AccountRepository {
	return &SingleFileAccountRepo{
		db: db,
	}
}

func (r *SingleFileAccountRepo) Insert(tx common.Tx, a *account.Account) (int32, error) {
	accounts := r.FindAll()
	accounts = append(accounts, a)
	r.db.SaveTable(ACCOUNT, func() []any {
		anyAccounts := make([]any, len(accounts))
		for i, acc := range accounts {
			anyAccounts[i] = acc
		}
		return anyAccounts
	}())
	return int32(1), nil
}

func (r *SingleFileAccountRepo) Update(tx common.Tx, a *account.Account) (int32, error) {
	panic("not implemented") // TODO: Implement
}
func (r *SingleFileAccountRepo) Save(tx common.Tx, a *account.Account) (id int32) {
	panic("not implemented") // TODO: Implement
}
func (r *SingleFileAccountRepo) Get(name string) (*account.Account, error) {
	panic("not implemented") // TODO: Implement
}
func (r *SingleFileAccountRepo) Find(name string) *account.Account {
	panic("not implemented") // TODO: Implement
}
func (r *SingleFileAccountRepo) FindAll() []*account.Account {
	return singlefile.ConvertToTypedSlide(r.db.GetTable(ACCOUNT), account.Account{})
}
