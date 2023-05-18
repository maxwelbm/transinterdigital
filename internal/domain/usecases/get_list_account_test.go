package usecases

import (
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"reflect"
	"testing"
)

func Test_useCase_GetListAccount(t *testing.T) {
	type fields struct {
		repository Repository
	}
	tests := []struct {
		name     string
		fields   fields
		want     []AccountOutput
		wantResp *helper.Response
	}{
		{
			name: "get list account with success",
			fields: fields{
				repository: Repository{
					account: entity.AccontMock{},
				},
			},
			want: []AccountOutput{
				{ID: 1, Name: "Max", CPF: "12312312312", Secret: "maxsecret", Balance: 1000000.12},
				{ID: 2, Name: "Salty", CPF: "32132132132", Secret: "saltysecret", Balance: 1000000.13},
			},
			wantResp: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}
			got, got1 := c.GetListAccount()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListAccount() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantResp) {
				t.Errorf("GetListAccount() got1 = %v, want %v", got1, tt.wantResp)
			}
		})
	}
}
