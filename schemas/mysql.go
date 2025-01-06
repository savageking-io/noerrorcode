package schemas

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type FieldType string

const (
	INT   FieldType = "INT"
	STR   FieldType = "STR"
	FLOAT FieldType = "FLOAT"
)

func (t *FieldType) Scan(value interface{}) error {
	*t = FieldType(value.([]byte))
	return nil
}

func (t FieldType) Value() (driver.Value, error) {
	return string(t), nil
}

type User struct {
	gorm.Model

	Username string
	Password string
	Email    string
	Steam    SteamUser
}

type SteamUser struct {
	gorm.Model

	UserID       uint
	SteamID      string
	OwnerSteamID string
	VAC          bool
	Ban          bool // Publisher ban
}

type Session struct {
	gorm.Model

	UID       uint
	IP        string
	UserAgent string
	Revision  string
	User      User `gorm:"foreignKey:UID"`
}

// CharacterStats belong to player characters
// Futher configuration can be made for specific type of stat in CharacterStatsInteger, CharacterStatsString and CharacterStatsFloat
// Stats are stored in MongoDB within characters collection
type CharacterStats struct {
	gorm.Model

	Name             string                // Name of the stat that will be returned to the game-client
	Description      string                // Small internal description for the stat
	ValueType        FieldType             `sql:"type:ENUM('INT', 'STR', 'FLOAT')" gorm:"column:value_type"` // Data type
	CanBeEmpty       bool                  // For string type stats
	IsRequired       bool                  // Whether or not this stat must be set when creating a new character
	IsUnique         bool                  // Whether or not this stat must be unique
	UniqueScope      uint                  // 1 - Global unique among all characters, 2 - Unique within user account
	UpdatePermission uint                  // 0 - Updated by the owner, 1 - can be updated by game server
	String           CharacterStatsString  // For string stats
	Integer          CharacterStatsInteger // For integer stats
	Float            CharacterStatsFloat   // For float stats
}

type CharacterStatsInteger struct {
	gorm.Model

	CharacterStatsID uint
	Default          int
	Min              int
	Max              int
}

type CharacterStatsString struct {
	gorm.Model

	CharacterStatsID uint
	Default          string
	Min              int
	Max              int
}

type CharacterStatsFloat struct {
	gorm.Model

	CharacterStatsID uint
	Default          float64
	Min              float64
	Max              float64
}
