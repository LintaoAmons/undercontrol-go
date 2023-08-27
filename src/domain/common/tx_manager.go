package common

type TxManager interface {
	Begin() Tx
}

type Tx interface {
	Commit() error
	Rollback() error
}

type DefaultTxManager struct{}

type DefaultTx struct{}

func (tm DefaultTxManager) Begin() Tx {
	return &DefaultTx{}
}

func (t DefaultTx) Commit() error {
	return nil
}

func (t DefaultTx) Rollback() error {
	return nil
}
