package main

import (
	"testing"

	"github.com/savageking-io/noerrorcode/database"
)

func TestConfig_Read(t *testing.T) {
	type fields struct {
		WebSocket *WebSocketConfig
		MySQL     *database.MySQLConfig
		MongoDB   *database.MongoDBConfig
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"No filepath", fields{}, args{}, true},
		{"Broken config", fields{}, args{filepath: "testdata/config/broken.yaml"}, true},
		{"No crypto", fields{}, args{filepath: "testdata/config/normal.yaml"}, true},
		{"Broken cryptography", fields{}, args{filepath: "testdata/config/broken-crypto.yaml"}, true},
		{"Working config", fields{}, args{filepath: "testdata/config/working.yaml"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				WebSocket: tt.fields.WebSocket,
				MySQL:     tt.fields.MySQL,
				MongoDB:   tt.fields.MongoDB,
			}
			if err := c.Read(tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("Config.Read() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
