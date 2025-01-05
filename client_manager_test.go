package main

import (
	"reflect"
	"sync"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/schemas"
	"github.com/savageking-io/noerrorcode/steam"
	_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func TestClientManager_GetUserBySteamID(t *testing.T) {
	conn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer conn.Close()

	//mock.ExpectExec("SELECT").WillReturnResult(sqlmock.NewResult(1, 1))
	rows0 := sqlmock.NewRows([]string{"id"})
	rows1 := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("SELECT").WillReturnRows(rows0)
	mock.ExpectQuery("SELECT").WillReturnRows(rows1)

	db := new(database.MySQL)

	dialector := _mysql.New(_mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      conn,
		SkipInitializeWithVersion: true,
	})
	db.DB, _ = gorm.Open(dialector, &gorm.Config{})

	user := &schemas.User{}
	user.ID = 1

	type fields struct {
		clients      map[uuid.UUID]*Client
		mutex        sync.Mutex
		steam        *steam.Steam
		cryptoConfig *CryptoConfig
		mysql        *database.MySQL
		mongo        *database.MongoDB
	}
	type args struct {
		steamID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *schemas.User
		wantErr bool
	}{
		{"Nil MySQL", fields{}, args{}, nil, true},
		{"Empty SteamID", fields{mysql: db}, args{}, nil, true},
		{"Not Found by SteamID", fields{mysql: db}, args{"000000000000000"}, nil, false},
		{"Found by SteamID", fields{mysql: db}, args{"000000000000000"}, user, false},
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
			got, err := d.GetUserBySteamID(tt.args.steamID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientManager.GetUserBySteamID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientManager.GetUserBySteamID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientManager_CreateUserFromSteam(t *testing.T) {
	type fields struct {
		clients      map[uuid.UUID]*Client
		mutex        sync.Mutex
		steam        *steam.Steam
		cryptoConfig *CryptoConfig
		mysql        *database.MySQL
		mongo        *database.MongoDB
	}
	type args struct {
		steamID      string
		ownerSteamID string
		vac          bool
		ban          bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *schemas.User
		wantErr bool
	}{
		{"Nil MySQL", fields{}, args{}, nil, true},
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
			got, err := d.CreateUserFromSteam(tt.args.steamID, tt.args.ownerSteamID, tt.args.vac, tt.args.ban)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientManager.CreateUserFromSteam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientManager.CreateUserFromSteam() = %v, want %v", got, tt.want)
			}
		})
	}
}
