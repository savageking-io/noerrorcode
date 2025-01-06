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
	MsgHeaderSize       uint32 = 8
	MsgTypeHello        uint32 = 0xFEF0BAB0
	MsgTypeWelcome      uint32 = 0xFEF0BAB1
	MsgTypeAuth         uint32 = 0xFEF0FEFA
	MsgTypeAuthResponse uint32 = 0xFEF0FEFB

	// Characters
	MsgTypeCharactersGet          uint32 = 0xC4A97157 // Get list of all player's character IDs
	MsgTypeCharacterGet           uint32 = 0xC4A97158 // Get Character by ID
	MsgTypeCharacterCreate        uint32 = 0xC4A97159 // Create new character
	MsgTypeCharacterValidateState uint32 = 0xC4A97160 // Validate specific stat
	MsgTypeCharacterSetStat       uint32 = 0xC4A97161 // Update stat value
	MsgTypeCharacterGetStat       uint32 = 0xC4A97162 // Get stat value
)

// Auth Status Codes 100XX
const (
	StatusCodeAuthSuccess         uint16 = 0     // Authentication succeed
	StatusCodeAuthInternalError   uint16 = 10001 // Internal server error (e.g. failed to unmarshal)
	StatusCodeAuthExternalError   uint16 = 10002 // External server error (platform-side)
	StatusCodeAuthAuthFailed      uint16 = 10003 // Authentication failed
	StatusCodeGenerateTokenFailed uint16 = 10004
)
