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
	MsgTypeAuth    uint32 = 0xFEF0FEFA
)

// Auth Status Codes 100XX
const (
	StatusCodeAuthSuccess         uint16 = 0     // Authentication succeed
	StatusCodeAuthInternalError   uint16 = 10001 // Internal server error (e.g. failed to unmarshal)
	StatusCodeAuthExternalError   uint16 = 10002 // External server error (platform-side)
	StatusCodeAuthAuthFailed      uint16 = 10003 // Authentication failed
	StatusCodeGenerateTokenFailed uint16 = 10004
)
