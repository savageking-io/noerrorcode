package main

var (
	AppVersion               string = "Undefined"
	ConfigFilepath           string = ""
	LogLevel                 string = "info"
	DefaultWebSocketURL      string = "/ws"
	TargetHostname           string = "localhost:31312"
	WebSocketReadBufferSize  int64  = 1024
	WebSocketWriteBufferSize int64  = 1024
	WebSocketPingTimeout     int    = 30
)

const (
	MsgTypeHello   uint32 = 0xFEF0BAB0
	MsgTypeWelcome uint32 = 0xFEF0BAB1
)
