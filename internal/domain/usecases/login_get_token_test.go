package usecases

import (
	"errors"
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"net/http"
	"reflect"
	"testing"
)

func Test_useCase_LoginGetToken(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		input TokenInput
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     Token
		wantResp *helper.Response
	}{
		{
			name: "validated flow get token",
			fields: fields{
				repository: Repository{
					account:  entity.AccontMock{},
					transfer: entity.TransfersMock{},
				},
			},
			args: args{
				input: TokenInput{
					CPF:       "415.333.420-12",
					Secret:    "supersecret",
					KeySecret: "keysecret",
				},
			},
			wantResp: &helper.Response{Status: http.StatusBadRequest, Err: errors.New("cpf invalid")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}

			got, err := c.LoginGetToken(tt.args.input)
			if !reflect.DeepEqual(err, tt.wantResp) {
				t.Errorf("LoginGetToken() err = %v, want %v", err, tt.wantResp)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginGetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
