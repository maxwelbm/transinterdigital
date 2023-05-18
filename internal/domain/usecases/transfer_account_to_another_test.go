package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"reflect"
	"testing"
)

type AccountMockTransferAccountToAnother struct{}
type TransfersMockTransferAccountToAnother struct{}

func (u AccountMockTransferAccountToAnother) Balance(accountID int) (float64, error) {
	return 100, nil
}

func (u AccountMockTransferAccountToAnother) List() ([]entity.Account, error) {
	return []entity.Account{
		{ID: 1, Name: "Max", CPF: "12312312312", Secret: "maxsecret", Balance: 1000000.12},
		{ID: 2, Name: "Salty", CPF: "32132132132", Secret: "saltysecret", Balance: 1000000.13},
	}, nil
}

func (u AccountMockTransferAccountToAnother) UpdateBalance(accountID int, balance float64) error {
	return nil
}

func (u AccountMockTransferAccountToAnother) GetAccountID(cpf, secret string) (int64, error) {
	return 0, nil
}

func (u AccountMockTransferAccountToAnother) Save(account *entity.Account) error {
	if len(account.Name) == 0 {
		return errors.New("failed in create account")
	}
	return nil
}

func (t TransfersMockTransferAccountToAnother) List(originID int64) ([]entity.Transfers, error) {
	if originID == 1 {
		return []entity.Transfers{{ID: 1, AccountDestinationID: 2, AccountOriginID: 1, Amount: 2}}, nil
	}
	return nil, nil
}

func (t TransfersMockTransferAccountToAnother) Save(transfers entity.Transfers) error {
	return nil
}

func Test_useCase_TransferAccountToAnother(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		input TransferInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *helper.Response
	}{
		{
			name: "transfer full",
			fields: fields{
				repository: Repository{
					account:  AccountMockTransferAccountToAnother{},
					transfer: TransfersMockTransferAccountToAnother{},
				},
			},
			args: args{
				input: TransferInput{1, 2, 12.2},
			},
			want: nil,
		},
		{
			name: "transfer to me",
			fields: fields{
				repository: Repository{
					account:  AccountMockTransferAccountToAnother{},
					transfer: TransfersMockTransferAccountToAnother{},
				},
			},
			args: args{
				input: TransferInput{1, 1, 12.2},
			},
			want: &helper.Response{Status: http.StatusBadRequest, Err: errors.New("there is no way to transfer it to yourself")},
		},
		{
			name: "transfer a little more",
			fields: fields{
				repository: Repository{
					account:  AccountMockTransferAccountToAnother{},
					transfer: TransfersMockTransferAccountToAnother{},
				},
			},
			args: args{
				input: TransferInput{1, 2, 111111111111.211},
			},
			want: &helper.Response{Status: http.StatusBadRequest, Err: errors.New("insufficient balance to complete the transfer")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}
			if got := c.TransferAccountToAnother(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransferAccountToAnother() = %v, want %v", got, tt.want)
			}
		})
	}
}
