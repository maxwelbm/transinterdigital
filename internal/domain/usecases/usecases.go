package usecases

import "github.com/maxwelbm/transinterdigital/internal/domain/entity"

type useCase struct {
	repository Repository
}

type Repository struct {
	account  entity.AccountRepository
	transfer entity.TransfersRepository
}

type UseCase interface {
	CreateAccount(input AccountInput) error
	GetListAccount() ([]AccountOutput, error)
	GetBalance(accountID int) (BalanceOutput, error)
	GetListTransfers(originID int64) ([]TransfersOutput, error)
	TransferAccountToAnother(input TransferInput) error
	LoginGetToken(input TokenInput) (Token, error)
}

func New(repository Repository) *useCase {
	return &useCase{
		repository: repository,
	}
}

func (u *Repository) SetAccount(account entity.AccountRepository) {
	u.account = account
}

func (u *Repository) SetTransfer(transfer entity.TransfersRepository) {
	u.transfer = transfer
}
