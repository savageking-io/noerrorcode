package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/savageking-io/noerrorcode/schemas"
	log "github.com/sirupsen/logrus"
)

type MessageSchema struct {
	Message string `json:"message"`
}

type Client struct {
	conn           *websocket.Conn
	uuid           uuid.UUID
	token          string
	PlatformUserID string
}

func (c *Client) Run() {
	log.Traceln("Client::Run")

	if c.conn == nil {
		log.Errorf("Client [%s]: nil conn", c.uuid)
		return
	}

	//c.conn.SetReadLimit(WebSocketReadBufferSize)
	//c.conn.SetReadDeadline(time.Now().Add(time.Duration(WebSocketPingTimeout)))

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

	ctrl := binary.BigEndian.Uint32(payload[:4])
	messageId := binary.BigEndian.Uint32(payload[4:8])
	switch ctrl {
	case MsgTypeHello:
		return c.HandleHello(messageId, payload[8:])
	case MsgTypeAuth:
		return c.HandleAuth(messageId, payload[8:])
	}

	return fmt.Errorf("bad packet %+v [%x]", payload, payload[:4])
}

func (c *Client) HandleHello(messageId uint32, data []byte) error {
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

	return c.Send(MsgTypeWelcome, messageId, welcome)
}

func (c *Client) HandleAuth(messageId uint32, payload []byte) error {
	log.Traceln("Client::HandleAuth")
	log.Debugf("Client [%s]: Requested auth", c.uuid)

	response := new(schemas.AuthResponse)

	if nec == nil {
		return fmt.Errorf("nil nec")
	}

	if nec.Steam == nil {
		response.Status = StatusCodeAuthInternalError
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("nil steam")
	}

	packet := new(schemas.AuthRequest)
	err := json.Unmarshal(payload, packet)
	if err != nil {
		response.Status = StatusCodeAuthInternalError
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("unmarshal failed. Client: %s, Data: %+v", c.uuid, payload)
	}

	log.Debugf("Client [%s]: Auth ticket: %s", c.uuid, packet.Ticket)

	steamResponse, err := nec.Steam.AuthUserTicket([]byte(packet.Ticket))
	if err != nil {
		response.Status = StatusCodeAuthExternalError
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("auth failed: %s", err.Error())
	}

	if steamResponse == nil || steamResponse.Response == nil || steamResponse.Response.Params == nil {
		response.Status = StatusCodeAuthInternalError
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("auth failed internally: nil data")
	}

	if steamResponse.Response.Params.Result != "OK" {
		response.Status = StatusCodeAuthAuthFailed
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("auth failed: %s", steamResponse.Response.Params.Result)
	}

	// Authentication succeed
	c.PlatformUserID = steamResponse.Response.Params.SteamID
	if err := c.GenerateToken(); err != nil {
		response.Status = StatusCodeGenerateTokenFailed
		c.Send(MsgTypeAuthResponse, messageId, response)
		return fmt.Errorf("token failed: %s", err.Error())
	}

	response.Status = 0
	response.Token = c.token
	return c.Send(MsgTypeAuthResponse, messageId, response)
}

func (c *Client) Send(msgType uint32, messageId uint32, v any) error {
	log.Traceln("Client::Send")
	data, err := json.Marshal(v)
	if err != nil {
		log.Debugf("Failed to marshal: %+v", v)
		return fmt.Errorf("marshal failed: %s", err.Error())
	}
	return c.SendRaw(c.MakeMessage(msgType, messageId, data))
}

func (c *Client) SendRaw(payload []byte) error {
	log.Traceln("Client::SendRaw")
	if c.conn == nil {
		return fmt.Errorf("nil connection")
	}
	return c.conn.WriteMessage(1, payload)
}
func (c *Client) MakeMessage(msgType uint32, messageId uint32, payload []byte) []byte {
	log.Traceln("Client::MakeMessage")
	var messageTypeHeader = make([]byte, 4)
	var messageIdHeader = make([]byte, 4)
	binary.BigEndian.PutUint32(messageTypeHeader, msgType)
	binary.BigEndian.PutUint32(messageIdHeader, messageId)
	return append(messageTypeHeader, append(messageIdHeader, payload...)...)
}

func (c *Client) GenerateToken() error {
	log.Traceln("User::CreateToken")

	if nec == nil {
		return fmt.Errorf("nil nec")
	}

	if nec.Config == nil {
		return fmt.Errorf("nil config")
	}

	if nec.Config.Crypto == nil {
		return fmt.Errorf("nil crypto config")
	}

	if nec.Config.Crypto.Key == "" {
		return fmt.Errorf("empty sign key")
	}

	var err error
	token, err := jwt.NewBuilder().
		Issuer("savageking.io"). // @TODO: Move to configuration
		IssuedAt(time.Now()).
		Claim("uid", c.PlatformUserID).
		Build()

	if err != nil {
		log.Errorf("Failed to build new token: %s", err.Error())
		return fmt.Errorf("Token build failed: %s", err.Error())
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(nec.Config.Crypto.Key)))
	if err != nil {
		log.Errorf("Failed to sign token: %s", err.Error())
		return fmt.Errorf("Token sign failed: %s", err.Error())
	}

	c.token = string(signed)
	log.Debugf("Client [%s] generated token: %s", c.uuid, c.token)
	return nil
}
