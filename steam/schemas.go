package steam

type AuthTicketResponse struct {
	Response *AuthTicketResponseParams `json:"response"`
}

type AuthTicketResponseParams struct {
	Params *AuthTicketResponsePayload `json:"params"`
}

type AuthTicketResponsePayload struct {
	Result          string `json:"result"`
	SteamID         string `json:"steamid"`
	OwnerSteamID    string `json:"ownersteamid"`
	VACBanned       bool   `json:"vacbanned"`
	PublisherBanned bool   `json:"publisherbanned"`
}

type UserTicketAuthRequest struct {
	Key      string `schema:"key"`
	AppId    uint32 `schema:"appid"`
	Ticket   string `schema:"ticket"`
	Identity string `schema:"identity"`
}
