package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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
