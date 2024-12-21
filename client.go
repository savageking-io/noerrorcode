package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn *websocket.Conn
	uuid uuid.UUID
}

func (c *Client) Run() {
	log.Traceln("Client::Run")

	c.conn.SetReadLimit(WebSocketReadBufferSize)
	c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))
	c.conn.SetPongHandler(c.PongHandler)

	for {
		n, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Client: read error [%s]: %v", c.uuid.String(), err.Error())
			}
			break
		}
		log.Debugf("Client: Read %d bytes: %s", n, string(msg))
	}
}

func (c *Client) Send(payload []byte) error {
	return c.conn.WriteMessage(1, payload)
}

func (c *Client) PongHandler(in string) error {
	c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))
	return nil
}
