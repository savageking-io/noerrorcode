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
}

type Session struct {
	gorm.Model

	UID       uint
	IP        string
	UserAgent string
	Revision  string
	User      User `gorm:"foreignKey:UID"`
}

type SteamLink struct {
	gorm.Model

	UID          uint
	User         User `gorm:"foreignKey:UID"`
	SteamID      string
	OwnerSteamID string
	VAC          bool
	Ban          bool // Publisher ban
}

type CharacterStats struct {
	gorm.Model

	Name        string
	Description string
	ValueType   FieldType `sql:"type:ENUM('INT', 'STR', 'FLOAT')" gorm:"column:value_type"`
	CanBeEmpty  bool
	String      CharacterStatsString
	Integer     CharacterStatsInteger
	Float       CharacterStatsFloat
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
	Default          float32
	Min              float32
	Max              float32
}
