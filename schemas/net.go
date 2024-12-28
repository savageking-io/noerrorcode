package schemas

type HelloMessage struct {
	Platform  string `json:"platform"`
	Revision  uint16 `json:"revision"`
	OSVersion string `json:"osVersion"`
}

type WelcomeMessage struct {
	Revision uint16 `json:"revision"`
	Version  string `json:"version"`
	Status   uint16 `json:"status"`
}

type AuthRequest struct {
	Ticket string `json:"ticket"`
}

type AuthResponse struct {
	Status uint16 `json:"status"`
	Token  string `json:"token"`
}
