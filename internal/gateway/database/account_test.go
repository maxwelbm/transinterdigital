package database

import (
	"github.com/maxwelbm/transinterdigital/internal/domain/entity"
	"testing"
)

func TestAccountRepository_Save(t *testing.T) {
	type fields struct {
		Db Conn
	}
	type args struct {
		account *entity.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "flow save success",
			fields:  fields{DbMock{}},
			args:    args{account: &entity.Account{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &AccountRepository{
				Db: tt.fields.Db,
			}
			if err := d.Save(tt.args.account); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
