package main

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func TestClient_Run(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	tests := []struct {
		name   string
		fields fields
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			c.Run()
		})
	}
}

func TestClient_Handle(t *testing.T) {
	p0 := []byte{}
	p1 := []byte{}
	pbad := []byte{}

	p0_type := make([]byte, 4)
	p0_msg_id := make([]byte, 4)
	binary.BigEndian.PutUint32(p0_type, MsgTypeHello)
	binary.BigEndian.PutUint32(p0_msg_id, 0)
	p0 = append(p0, p0_type...)
	p0 = append(p0, p0_msg_id...)
	p0 = append(p0, []byte("{}")...)

	p1_type := make([]byte, 4)
	p1_msg_id := make([]byte, 4)
	binary.BigEndian.PutUint32(p1_type, MsgTypeAuth)
	binary.BigEndian.PutUint32(p1_msg_id, 0)
	p1 = append(p1, p1_type...)
	p1 = append(p1, p1_msg_id...)
	p1 = append(p1, []byte("{}")...)

	pbad_type := make([]byte, 4)
	binary.BigEndian.PutUint32(pbad_type, 0)
	pbad = append(pbad, pbad_type...)

	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Empty payload", fields{}, args{}, true},
		{"Hello packet type", fields{}, args{payload: p0}, true},
		{"Auth packet type", fields{}, args{payload: p1}, true},
		{"Bad packet type", fields{}, args{payload: pbad}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.Handle(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Client.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_HandleHello(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		messageId uint32
		data      []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Bad payload", fields{}, args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.HandleHello(tt.args.messageId, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Client.HandleHello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_HandleAuth(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		messageId uint32
		payload   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.HandleAuth(tt.args.messageId, tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Client.HandleAuth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		msgType   uint32
		messageId uint32
		v         any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.Send(tt.args.msgType, tt.args.messageId, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SendRaw(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.SendRaw(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Client.SendRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_MakeMessage(t *testing.T) {
	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	type args struct {
		msgType   uint32
		messageId uint32
		payload   []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if got := c.MakeMessage(tt.args.msgType, tt.args.messageId, tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.MakeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GenerateToken(t *testing.T) {
	//nec0 := new(NoErrorCode)

	type fields struct {
		conn           *websocket.Conn
		uuid           uuid.UUID
		token          string
		PlatformUserID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Nil Nec", fields{}, true},
		{"Nil Config", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn:           tt.fields.conn,
				uuid:           tt.fields.uuid,
				token:          tt.fields.token,
				PlatformUserID: tt.fields.PlatformUserID,
			}
			if err := c.GenerateToken(); (err != nil) != tt.wantErr {
				t.Errorf("Client.GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
