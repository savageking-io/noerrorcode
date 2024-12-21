package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  int(WebSocketReadBufferSize),
	WriteBufferSize: int(WebSocketWriteBufferSize),
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocket struct {
	config   *WebSocketConfig
	connchan chan *websocket.Conn
	clients  map[uuid.UUID]*Client
	mutex    sync.Mutex
}

func (ws *WebSocket) Init(config *WebSocketConfig) error {
	log.Traceln("WebSocket::Init")
	if config == nil {
		return fmt.Errorf("websocket: nil config")
	}
	ws.config = config
	if ws.config.URL == "" {
		ws.config.URL = DefaultWebSocketURL
	}
	ws.connchan = make(chan *websocket.Conn)
	ws.clients = make(map[uuid.UUID]*Client)
	return nil
}

func (ws *WebSocket) Run() error {
	log.Traceln("WebSocket::Run")
	log.Infof("Starting WebSocket server on %s:%d", ws.config.Hostname, ws.config.Port)
	http.HandleFunc(ws.config.URL, ws.Handle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", ws.config.Hostname, ws.config.Port), nil))
	return nil
}

func (ws *WebSocket) Handle(w http.ResponseWriter, r *http.Request) {
	log.Traceln("WebSocket::Handle")

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("WebSocket: failed to upgrade new incoming connectiom from %s: %s", r.RemoteAddr, err.Error())
		return
	}

	u := uuid.New()

	log.Debugf("New incoming connection upgraded from %s as %s", r.RemoteAddr, u)
	defer ws.UnregisterClient(u)

	if err := ws.RegisterClient(c, u); err != nil {
		log.Errorf("WebSocket: failed to register client: %s", err.Error())
		return
	}

}

func (ws *WebSocket) RegisterClient(conn *websocket.Conn, id uuid.UUID) error {
	log.Traceln("WebSocket::RegisterClient")
	if conn == nil {
		return fmt.Errorf("nil conn")
	}
	log.Infof("WebSocket: Registering new client %s [%s]", id, conn.RemoteAddr().String())
	client := &Client{
		conn: conn,
		uuid: id,
	}
	ws.mutex.Lock()
	ws.clients[id] = client
	ws.mutex.Unlock()
	ws.clients[id].Run()
	return nil
}

func (ws *WebSocket) UnregisterClient(id uuid.UUID) {
	log.Traceln("WebSocket::UnregisterClient")
	client, exists := ws.clients[id]
	if !exists {
		log.Debugf("WebSocket: failed to unregister client: %s not found", id.String())
		return
	}
	log.Infof("WebSocket: Unregistering client %s", id.String())
	ws.mutex.Lock()
	if client.conn != nil {
		client.conn.Close()
	}
	delete(ws.clients, id)
	ws.mutex.Unlock()
}
