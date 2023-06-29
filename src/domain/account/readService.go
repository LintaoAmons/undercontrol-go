package account

import "github.com/pterm/pterm"

// For more complix queries, easy ones can just done by readRepo
type AccountReadService struct {
	repo   AccountReadRepo
	logger *pterm.Logger
}

func NewAccountReadService(repo AccountReadRepo) *AccountReadService {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace).WithMaxWidth(200)
	return &AccountReadService{
		repo:   repo,
		logger: logger,
	}
}
