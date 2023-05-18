package usecases

import (
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"reflect"
	"testing"
)

func Test_useCase_GetBalance(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		accountID int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     BalanceOutput
		wantResp *helper.Response
	}{
		{
			name: "get balance with success",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				accountID: 1,
			},
			want:     BalanceOutput{1000000.12},
			wantResp: nil,
		},
		{
			name: "there are no items",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				accountID: 3,
			},
			want:     BalanceOutput{0},
			wantResp: &helper.Response{Status: http.StatusNotFound, Err: pgx.ErrNoRows},
		},
		{
			name: "failed in get item",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				accountID: 2,
			},
			want:     BalanceOutput{0},
			wantResp: &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in get balance")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}
			got, got1 := c.GetBalance(tt.args.accountID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantResp) {
				t.Errorf("GetBalance() got1 = %v, want %v", got1, tt.wantResp)
			}
		})
	}
}
