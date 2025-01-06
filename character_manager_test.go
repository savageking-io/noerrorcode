package main

import (
	"testing"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/schemas"
)

func TestCharacterManager_Init(t *testing.T) {
	mysql := new(database.MySQL)

	type fields struct {
		mysql *database.MySQL
		stats []schemas.CharacterStats
	}
	type args struct {
		mysql *database.MySQL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Nil MySQL", fields{}, args{}, true},
		{"Failed to load", fields{}, args{mysql}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &CharacterManager{
				mysql: tt.fields.mysql,
				stats: tt.fields.stats,
			}
			if err := d.Init(tt.args.mysql); (err != nil) != tt.wantErr {
				t.Errorf("CharacterManager.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
