package main

import (
	"encoding/binary"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	log "github.com/sirupsen/logrus"
)

func TestClient_Run(t *testing.T) {
	type fields struct {
		conn *websocket.Conn
		uuid uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn: tt.fields.conn,
				uuid: tt.fields.uuid,
			}
			c.Run()
		})
	}
}

func TestClient_Handle(t *testing.T) {
	log.SetLevel(log.TraceLevel)

	helloHeader := make([]byte, 4)
	binary.BigEndian.PutUint32(helloHeader, MsgTypeHello)

	emptyJson := "{}"
	helloMessage := helloHeader
	helloMessage = append(helloMessage, []byte(emptyJson)...)

	type fields struct {
		conn *websocket.Conn
		uuid uuid.UUID
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
		{"Small payload", fields{}, args{}, true},
		{"Hello Message", fields{}, args{helloMessage}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				conn: tt.fields.conn,
				uuid: tt.fields.uuid,
			}
			if err := c.Handle(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Client.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_HandleHello(t *testing.T) {
	type fields struct {
		conn *websocket.Conn
		uuid uuid.UUID
	}
	type args struct {
		data []byte
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
				conn: tt.fields.conn,
				uuid: tt.fields.uuid,
			}
			if err := c.HandleHello(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Client.HandleHello() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Send(t *testing.T) {
	type fields struct {
		conn *websocket.Conn
		uuid uuid.UUID
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
				conn: tt.fields.conn,
				uuid: tt.fields.uuid,
			}
			if err := c.Send(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("Client.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_PongHandler(t *testing.T) {
	type fields struct {
		conn *websocket.Conn
		uuid uuid.UUID
	}
	type args struct {
		in string
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
				conn: tt.fields.conn,
				uuid: tt.fields.uuid,
			}
			if err := c.PongHandler(tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("Client.PongHandler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
