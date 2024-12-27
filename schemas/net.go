package schemas

type HelloMessage struct {
	Platform  string `json:"platform"`
	Revision  uint16 `json:"revision"`
	OSVersion string `json:"osVersion"`
}
