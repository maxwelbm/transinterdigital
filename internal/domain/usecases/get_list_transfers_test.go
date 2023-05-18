package usecases

import (
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"github.com/maxwelbm/transinterdigital/pkg/helper"
	"reflect"
	"testing"
)

func Test_useCase_GetListTransfers(t *testing.T) {
	type fields struct {
		repository Repository
	}
	type args struct {
		originID int64
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []TransfersOutput
		wantResp *helper.Response
	}{
		{
			name: "transfer full success",
			fields: fields{
				repository: Repository{
					account:  entity.AccontMock{},
					transfer: entity.TransfersMock{},
				},
			},
			args:     args{1},
			want:     []TransfersOutput{{ID: 1, AccountDestinationID: 2, AccountOriginID: 1, Amount: 2}},
			wantResp: nil,
		},
		{
			name: "list empty",
			fields: fields{
				repository: Repository{
					account:  entity.AccontMock{},
					transfer: entity.TransfersMock{},
				},
			},
			args:     args{32},
			want:     []TransfersOutput{},
			wantResp: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &useCase{
				repository: tt.fields.repository,
			}
			got, got1 := c.GetListTransfers(tt.args.originID)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListTransfers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantResp) {
				t.Errorf("GetListTransfers() got1 = %v, want %v", got1, tt.wantResp)
			}
		})
	}
}
