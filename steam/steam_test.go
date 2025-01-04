package steam

import (
	"reflect"
	"testing"
)

func TestSteam_Init(t *testing.T) {
	c0 := new(Config)
	c1 := new(Config)
	c1.AppId = 1
	c2 := new(Config)
	c2.AppId = 2
	c2.PublisherId = "test"
	c3 := new(Config)
	c3.AppId = 3
	c3.PublisherId = "test"
	c3.Key = "test"

	type fields struct {
		config *Config
		URL    string
	}
	type args struct {
		config *Config
		url    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Nil config", fields{}, args{}, true},
		{"No AppId", fields{}, args{config: c0}, true},
		{"Empty Publisher ID", fields{}, args{config: c1}, true},
		{"Empty Key", fields{}, args{config: c2}, true},
		{"Empty URL", fields{}, args{config: c3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Steam{
				config: tt.fields.config,
				URL:    tt.fields.URL,
			}
			if err := d.Init(tt.args.config, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Steam.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSteam_AuthUserTicket(t *testing.T) {
	c0 := new(Config)
	c1 := new(Config)
	c1.AppId = 1
	c2 := new(Config)
	c2.AppId = 2
	c2.PublisherId = "test"
	c3 := new(Config)
	c3.AppId = 3
	c3.PublisherId = "test"
	c3.Key = "test"
	type fields struct {
		config *Config
		URL    string
	}
	type args struct {
		authTicket []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthTicketResponse
		wantErr bool
	}{
		{"Nil config", fields{URL: "localhost"}, args{}, nil, true},
		{"No AppId", fields{URL: "localhost", config: c0}, args{}, nil, true},
		{"No PublisherID", fields{URL: "localhost", config: c1}, args{}, nil, true},
		{"No Key", fields{URL: "localhost", config: c2}, args{}, nil, true},
		{"No Key", fields{URL: "localhost", config: c3}, args{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Steam{
				config: tt.fields.config,
				URL:    tt.fields.URL,
			}
			got, err := d.AuthUserTicket(tt.args.authTicket)
			if (err != nil) != tt.wantErr {
				t.Errorf("Steam.AuthUserTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Steam.AuthUserTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
