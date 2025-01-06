package main

import (
	"fmt"

	"github.com/savageking-io/noerrorcode/database"
	"github.com/savageking-io/noerrorcode/schemas"
	log "github.com/sirupsen/logrus"
)

type CharacterManager struct {
	mysql *database.MySQL
	stats []schemas.CharacterStats
}

func (d *CharacterManager) Init(mysql *database.MySQL) error {
	log.Traceln("CharacterManager::Init")
	if mysql == nil {
		return fmt.Errorf("nil mysql")
	}
	d.mysql = mysql
	if err := d.LoadCharacterStats(); err != nil {
		return err
	}
	return nil
}

func (d *CharacterManager) LoadCharacterStats() error {
	log.Traceln("CharacterManager::LoadCharacterStats")

	if d.mysql == nil {
		return fmt.Errorf("load character stats: nil mysql")
	}
	if d.mysql.Get() == nil {
		return fmt.Errorf("mysql not initialized")
	}

	// Load all character stats
	result := d.mysql.Get().Find(&d.stats)
	if result.Error != nil {
		return fmt.Errorf("load character stats: %s", result.Error.Error())
	}
	log.Infof("Character Manager: Loaded %d stats", result.RowsAffected)

	return nil
}
