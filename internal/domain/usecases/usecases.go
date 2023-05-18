package usecases

import (
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
)

type useCase struct {
	repository Repository
}

type Repository struct {
	account  entity.AccountRepository
	transfer entity.TransfersRepository
}

type UseCase interface {
	CreateAccount(input AccountInput) *helper.Response
	GetListAccount() ([]AccountOutput, *helper.Response)
	GetBalance(accountID int) (BalanceOutput, *helper.Response)
	GetListTransfers(originID int64) ([]TransfersOutput, *helper.Response)
	TransferAccountToAnother(input TransferInput) *helper.Response
	LoginGetToken(input TokenInput) (Token, *helper.Response)
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
