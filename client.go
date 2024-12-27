package main

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type MessageSchema struct {
	Message string `json:"message"`
}

type Client struct {
	conn *websocket.Conn
	uuid uuid.UUID
}

func (c *Client) Run() {
	log.Traceln("Client::Run")

	//c.conn.SetReadLimit(WebSocketReadBufferSize)
	//c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))
	//c.conn.SetPongHandler(c.PongHandler)

	log.Infof("Working with client %s [%s]", c.uuid, c.conn.RemoteAddr().String())

	for {
		n, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Client: read error [%s]: %v", c.uuid.String(), err.Error())
			}
			break
		}
		log.Debugf("Client: Read %d bytes: %s", n, string(msg))
		c.Handle(msg)
	}
}

func (c *Client) Handle(payload []byte) error {
	data := new(MessageSchema)
	err := json.Unmarshal(payload, data)
	if err != nil {
		log.Errorf("Client [%s]: Failed to parse: %s", c.uuid, err.Error())
		return err
	}

	if data.Message == "hello" {
		response := new(MessageSchema)
		response.Message = "welcome"
		out, err := json.Marshal(response)
		if err != nil {
			log.Errorf("Client [%s]: Failed to marshal: %s", c.uuid, err.Error())
			return err
		}
		c.Send(out)
	}

	return nil
}

func (c *Client) Send(payload []byte) error {
	return c.conn.WriteMessage(1, payload)
}

func (c *Client) PongHandler(in string) error {
	c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))
	return nil
}
