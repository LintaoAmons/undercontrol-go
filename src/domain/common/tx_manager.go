package common

type TxManager interface {
	Begin() Tx
}

type Tx interface {
	Commit() error
	Rollback() error
}
