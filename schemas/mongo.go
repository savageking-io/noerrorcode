package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuthLog struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Ticket          []byte             `bson:"ticket,omitempty"`
	Version         string             `bson:"version,omitempty"`
	RequestSteamID  string             `bson:"request_steam_id,omitempty"`
	SteamUsername   string             `bson:"steam_username,omitempty"`
	SteamAuthResult string             `bson:"steam_auth_result,omitempty"`
	SteamID         string             `bson:"steam_id,omitempty"`
	OwnerSteamID    string             `bson:"owner_steam_id,omitempty"`
	VACBanned       bool               `bson:"vac_banned,omitempty"`
	PublisherBanned bool               `bson:"publisher_banned,omitempty"`
	Date            time.Time          `bson:"date,omitempty"`
}

type UserSession struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Platform  string             `bson:"platform,omitempty"`
	PlayerID  string             `bson:"player_id,omitempty"` // Platform-specific player ID
	Token     string             `bson:"token,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type ServiceRequestLogEntry struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	IP        string             `bson:"ip,omitempty"`
	UserAgent string             `bson:"user_agent,omitempty"`
	Date      time.Time          `bson:"date,omitempty"`
	Version   string             `bson:"version,omitempty"`
}
