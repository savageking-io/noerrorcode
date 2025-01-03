package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  int(WebSocketReadBufferSize),
	WriteBufferSize: int(WebSocketWriteBufferSize),
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocket struct {
	config  *WebSocketConfig
	clients *ClientManager
}

func (ws *WebSocket) Init(config *WebSocketConfig, clientManager *ClientManager) error {
	log.Traceln("WebSocket::Init")
	if config == nil {
		return fmt.Errorf("websocket: nil config")
	}
	if clientManager == nil {
		return fmt.Errorf("websocket: nil client manager")
	}
	ws.clients = clientManager
	ws.config = config
	if ws.config.URL == "" {
		ws.config.URL = DefaultWebSocketURL
	}

	return nil
}

func (ws *WebSocket) Run() error {
	log.Traceln("WebSocket::Run")
	log.Infof("Starting WebSocket server on %s:%d%s", ws.config.Hostname, ws.config.Port, ws.config.URL)
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

	log.Debugf("New incoming connection from %s", r.RemoteAddr)

	client, err := ws.clients.RegisterClient(c)
	if err != nil {
		log.Errorf("Failed to register new client: %s", err.Error())
		if err := c.Close(); err != nil {
			log.Errorf("Failed to close broken connection: %s", err.Error())
		}
		return
	}

	defer ws.clients.UnregisterClient(client.uuid)
	client.Run(ws.clients)
}
