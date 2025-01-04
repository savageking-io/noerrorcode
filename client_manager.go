package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/schemas"
	"github.com/savageking-io/noerrorcode/steam"
	log "github.com/sirupsen/logrus"
)

type ClientManager struct {
	clients      map[uuid.UUID]*Client
	mutex        sync.Mutex
	steam        *steam.Steam
	cryptoConfig *CryptoConfig
	mysql        *database.MySQL
	mongo        *database.MongoDB
}

func (d *ClientManager) Init(steam *steam.Steam, mysql *database.MySQL, mongo *database.MongoDB, cryptoConfig *CryptoConfig) error {
	if steam == nil {
		return fmt.Errorf("client manager: nil steam")
	}
	if mysql == nil {
		return fmt.Errorf("client manager: nil mysql")
	}
	if mongo == nil {
		return fmt.Errorf("client manager: nil mongo")
	}
	if cryptoConfig == nil {
		return fmt.Errorf("client manager: nil crypto config")
	}
	d.steam = steam
	d.cryptoConfig = cryptoConfig
	d.mysql = mysql
	d.mongo = mongo
	d.clients = make(map[uuid.UUID]*Client)
	return nil
}

func (d *ClientManager) RegisterClient(conn *websocket.Conn) (*Client, error) {
	log.Traceln("ClientManager::RegisterClient")
	if conn == nil {
		return nil, fmt.Errorf("client manager failed to register new client: nil conn")
	}
	client := &Client{
		conn: conn,
		uuid: uuid.New(),
	}
	d.mutex.Lock()
	d.clients[client.uuid] = client
	d.mutex.Unlock()

	return client, nil
}

func (d *ClientManager) UnregisterClient(uuid uuid.UUID) error {
	log.Traceln("ClientManager::UnregisterClient")
	client, exists := d.clients[uuid]
	if !exists {
		return fmt.Errorf("failed to unregister: %s doesn't exists", uuid.String())
	}

	d.mutex.Lock()
	if client.conn != nil {
		client.conn.Close()
	}
	delete(d.clients, uuid)
	d.mutex.Unlock()
	return nil
}

func (d *ClientManager) GenerateToken(userID string) (string, error) {
	log.Traceln("User::CreateToken")

	if d.cryptoConfig == nil {
		return "", fmt.Errorf("nil crypto config")
	}

	if d.cryptoConfig.Key == "" {
		return "", fmt.Errorf("empty sign key")
	}

	if userID == "" {
		return "", fmt.Errorf("empty user id")
	}

	var err error
	token, err := jwt.NewBuilder().
		Issuer(d.cryptoConfig.Issuer).
		IssuedAt(time.Now()).
		Claim("uid", userID).
		Build()

	if err != nil {
		return "", fmt.Errorf("token build failed: %s", err.Error())
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.HS256(), []byte(d.cryptoConfig.Key)))
	if err != nil {
		return "", fmt.Errorf("token sign failed: %s", err.Error())
	}

	return string(signed), nil
}

func (d *ClientManager) GetUserBySteamID(steamID string) (*schemas.User, error) {
	log.Traceln("ClientManager::GetUserBySteamID")

	if d.mysql == nil {
		return nil, fmt.Errorf("nil mysql")
	}

	if steamID == "" {
		return nil, fmt.Errorf("empty steam id")
	}

	user := &schemas.User{}
	result := d.mysql.Get().Joins("Steam").Where("steam_users.user_id = ?", steamID).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("get user by steam ID: %s", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, nil
}

func (d *ClientManager) CreateUserFromSteam(steamID, ownerSteamID string, vac, ban bool) (*schemas.User, error) {
	log.Traceln("ClientManager::CreateUserFromSteam")

	if d.mysql == nil {
		return nil, fmt.Errorf("nil mysql")
	}
	if steamID == "" {
		return nil, fmt.Errorf("empty steam id")
	}
	if ownerSteamID == "" {
		return nil, fmt.Errorf("empty owner steam id")
	}

	user := &schemas.User{
		Steam: schemas.SteamUser{
			SteamID:      steamID,
			OwnerSteamID: ownerSteamID,
			VAC:          vac,
			Ban:          ban,
		},
	}

	result := d.mysql.Get().Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("create: %s", result.Error.Error())
	}
	if result.RowsAffected != 1 {
		return nil, fmt.Errorf("nothing created")
	}

	return user, nil
}
