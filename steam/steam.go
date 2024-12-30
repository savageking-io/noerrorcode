package steam

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const api_backend = "https://partner.steam-api.com"

type Config struct {
	PublisherId string `yaml:"publisher_id"`
	AppId       uint32 `yaml:"app_id"`
	Key         string `yaml:"key"`
}

type Steam struct {
	config *Config
}

func (d *Steam) Init(config *Config) error {
	log.Traceln("Steam::Init")
	if config == nil {
		return fmt.Errorf("nil config")
	}
	if config.AppId == 0 {
		return fmt.Errorf("bad app id")
	}
	if config.PublisherId == "" {
		return fmt.Errorf("bad publisher id")
	}
	if config.Key == "" {
		return fmt.Errorf("bad key")
	}
	log.Debugf("Steam App ID: %d", config.AppId)
	log.Debugf("Steam Publisher ID: %s", config.PublisherId)
	log.Debugf("Steam Key: %s", config.Key)
	d.config = config
	return nil
}

func (d *Steam) AuthUserTicket(authTicket []byte) (*AuthTicketResponse, error) {
	log.Traceln("Steam::AuthUserTicket")

	ticket := hex.EncodeToString(authTicket)

	data := &UserTicketAuthRequest{
		Key:      d.config.Key,
		AppId:    uint32(d.config.AppId),
		Ticket:   ticket,
		Identity: fmt.Sprintf("WebAPI:%s", d.config.PublisherId),
	}

	payload := fmt.Sprintf("key=%s&appid=%d&ticket=%s&identity=%s", data.Key, data.AppId, data.Ticket, data.Identity)
	url := fmt.Sprintf("%s%s?%s", api_backend, "/ISteamUserAuth/AuthenticateUserTicket/v1/", payload)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		log.Errorf("Error creating request: %s", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request send failed: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read failed: %s", err.Error())
	}

	log.Debugf("Response: %+v", resp)
	log.Debugf("Body: %s", string(body))

	response := new(AuthTicketResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %s", err.Error())
	}

	return response, nil
}
