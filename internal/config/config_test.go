package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name string
		want Config
	}{
		{
			name: "checking if loaded that variables",
			want: Config{
				DBDriver:               "DB_DRIVER",
				DBHost:                 "DB_HOST",
				DBPort:                 "DB_PORT",
				DBUser:                 "DB_USER",
				DBPassword:             "DB_PASSWORD",
				DBName:                 "DB_NAME",
				KeySecret:              "KEY_SECRET",
				DBURL:                  "DB_URL",
				PgAdminDefaultEmail:    "PGADMIN_DEFAULT_EMAIL",
				PgAdminDefaultPassword: "PGADMIN_DEFAULT_PASSWORD",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
