package main

import (
	"net/http"
	"sync"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func TestWebSocket_Init(t *testing.T) {
	type fields struct {
		config   *WebSocketConfig
		connchan chan *websocket.Conn
		clients  map[uuid.UUID]*Client
		mutex    sync.Mutex
	}
	type args struct {
		config *WebSocketConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"No Config", fields{}, args{}, true},
		{"Empty URL", fields{}, args{config: new(WebSocketConfig)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WebSocket{
				config:   tt.fields.config,
				connchan: tt.fields.connchan,
				clients:  tt.fields.clients,
				mutex:    tt.fields.mutex,
			}
			if err := ws.Init(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("WebSocket.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebSocket_Run(t *testing.T) {
	type fields struct {
		config   *WebSocketConfig
		connchan chan *websocket.Conn
		clients  map[uuid.UUID]*Client
		mutex    sync.Mutex
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WebSocket{
				config:   tt.fields.config,
				connchan: tt.fields.connchan,
				clients:  tt.fields.clients,
				mutex:    tt.fields.mutex,
			}
			if err := ws.Run(); (err != nil) != tt.wantErr {
				t.Errorf("WebSocket.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebSocket_Handle(t *testing.T) {
	type fields struct {
		config   *WebSocketConfig
		connchan chan *websocket.Conn
		clients  map[uuid.UUID]*Client
		mutex    sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WebSocket{
				config:   tt.fields.config,
				connchan: tt.fields.connchan,
				clients:  tt.fields.clients,
				mutex:    tt.fields.mutex,
			}
			ws.Handle(tt.args.w, tt.args.r)
		})
	}
}

func TestWebSocket_RegisterClient(t *testing.T) {
	type fields struct {
		config   *WebSocketConfig
		connchan chan *websocket.Conn
		clients  map[uuid.UUID]*Client
		mutex    sync.Mutex
	}
	type args struct {
		conn *websocket.Conn
		id   uuid.UUID
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
			ws := &WebSocket{
				config:   tt.fields.config,
				connchan: tt.fields.connchan,
				clients:  tt.fields.clients,
				mutex:    tt.fields.mutex,
			}
			if err := ws.RegisterClient(tt.args.conn, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("WebSocket.RegisterClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebSocket_UnregisterClient(t *testing.T) {
	type fields struct {
		config   *WebSocketConfig
		connchan chan *websocket.Conn
		clients  map[uuid.UUID]*Client
		mutex    sync.Mutex
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws := &WebSocket{
				config:   tt.fields.config,
				connchan: tt.fields.connchan,
				clients:  tt.fields.clients,
				mutex:    tt.fields.mutex,
			}
			ws.UnregisterClient(tt.args.id)
		})
	}
}
