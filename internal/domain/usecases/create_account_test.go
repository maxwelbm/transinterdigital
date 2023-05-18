package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"reflect"
	"testing"
)

func Test_useCase_CreateAccount(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		input AccountInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *helper.Response
	}{
		{
			name: "created the account with success",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				input: AccountInput{
					Name:    "Max",
					CPF:     "573.835.700-06",
					Secret:  "superset",
					Balance: 1231.123,
				},
			},
			want: nil,
		},
		{
			name: "successfully created account without special character",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				input: AccountInput{
					Name:    "Max",
					CPF:     "57383570006",
					Secret:  "superset",
					Balance: 1231.123,
				},
			},
			want: nil,
		},
		{
			name: "created with cpf invalid",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				input: AccountInput{
					Name:    "Max",
					CPF:     "1232312312",
					Secret:  "superset",
					Balance: 1231.123,
				},
			},
			want: &helper.Response{Status: http.StatusBadRequest, Err: errors.New("cpf invalid")},
		},
		{
			name: "created with cpf invalid",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				input: AccountInput{
					Name:    "Max",
					CPF:     "1232312312",
					Secret:  "superset",
					Balance: 1231.123,
				},
			},
			want: &helper.Response{Status: http.StatusBadRequest, Err: errors.New("cpf invalid")},
		},
		{
			name: "failed created account",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			args: args{
				input: AccountInput{
					CPF:     "57383570006",
					Secret:  "superset",
					Balance: 1231.123,
				},
			},
			want: &helper.Response{Status: http.StatusInternalServerError, Err: errors.New("failed in save account")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}
			if got := c.CreateAccount(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
