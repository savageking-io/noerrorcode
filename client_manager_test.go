package main

import (
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/steam"
)

func TestClientManager_GenerateToken(t *testing.T) {
	c0 := new(CryptoConfig)
	c1 := new(CryptoConfig)
	c1.Key = "test"

	type fields struct {
		clients      map[uuid.UUID]*Client
		mutex        sync.Mutex
		steam        *steam.Steam
		cryptoConfig *CryptoConfig
		mysql        *database.MySQL
		mongo        *database.MongoDB
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"Nil CryptoConfig", fields{}, args{}, "", true},
		{"Empty key", fields{cryptoConfig: c0}, args{}, "", true},
		{"No User ID", fields{cryptoConfig: c1}, args{}, "", true},
		{"Passing", fields{cryptoConfig: c1}, args{userID: "111"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &ClientManager{
				clients:      tt.fields.clients,
				mutex:        tt.fields.mutex,
				steam:        tt.fields.steam,
				cryptoConfig: tt.fields.cryptoConfig,
				mysql:        tt.fields.mysql,
				mongo:        tt.fields.mongo,
			}
			got, err := d.GenerateToken(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientManager.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(tt.want) > 0 && got != tt.want {
				t.Errorf("ClientManager.GenerateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
