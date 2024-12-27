package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/savageking-io/noerrorcode/schemas"
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
		if err := c.Handle(msg); err != nil {
			log.Errorf("Client %s: %s", c.uuid, err.Error())
		}
	}
}

func (c *Client) Handle(payload []byte) error {
	if len(payload) < 4 {
		return fmt.Errorf("payload is too small: %d bytes", len(payload))
	}

	//ctrl := binary.BigEndian.Uint32(payload[:4])
	ctrl := binary.BigEndian.Uint32(payload[:4])
	switch ctrl {
	case MsgTypeHello:
		return c.HandleHello(payload[4:])
	}

	/*
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
	*/

	return fmt.Errorf("bad packet %+v [%x]", payload, payload[:4])
}

func (c *Client) HandleHello(data []byte) error {
	log.Traceln("Client::HandleHello")
	log.Debugf("Client [%s]: Received Hello", c.uuid)
	packet := new(schemas.HelloMessage)
	err := json.Unmarshal(data, packet)
	if err != nil {
		return fmt.Errorf("unmarshal failed. Client: %s, Data: %+v", c.uuid, data)
	}

	log.Debugf("Client [%s] is welcome. Revision: %d, OS: %s", c.uuid, packet.Revision, packet.OSVersion)

	welcome := new(schemas.WelcomeMessage)
	welcome.Revision = packet.Revision
	welcome.Status = 0 // @TODO: May be different if status not operational
	welcome.Version = AppVersion

	payload, err := json.Marshal(welcome)
	if err != nil {
		return fmt.Errorf("marshal failed: %s", err.Error())
	}

	return c.Send(c.MakeMessage(MsgTypeWelcome, payload))
}

func (c *Client) Send(payload []byte) error {
	if c.conn == nil {
		return fmt.Errorf("nil connection")
	}
	return c.conn.WriteMessage(1, payload)
}

func (c *Client) PongHandler(in string) error {
	c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))
	return nil
}

func (c *Client) MakeMessage(msgType uint32, payload []byte) []byte {
	var header = make([]byte, 4)
	binary.BigEndian.PutUint32(header, msgType)
	return append(header, payload...)
}
