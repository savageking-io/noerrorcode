package database

import (
	"fmt"

	_mysql "github.com/go-sql-driver/mysql"
	"github.com/savageking-io/noerrorcode/schemas"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Hostname      string `yaml:"hostname"`
	Port          uint16 `yaml:"port"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Database      string `yaml:"database"`
	Retry         bool   `yaml:"retry"`
	RetryAttempts int    `yaml:"retry_attempts"`
	RetryTimeout  int    `yaml:"retry_timeout"`
}

type MySQL struct {
	config *MySQLConfig
	db     *gorm.DB
}

func (db *MySQL) Init(config *MySQLConfig) error {
	log.Traceln("MySQL::Init")
	if config == nil {
		return fmt.Errorf("mysql init: nil config")
	}
	db.config = config
	return nil
}

func (db *MySQL) Connect() error {
	log.Traceln("MySQL::Connect")

	if db.config == nil {
		return fmt.Errorf("mysql: nil config")
	}

	conf := &_mysql.Config{
		User:      db.config.Username,
		Passwd:    db.config.Password,
		Addr:      fmt.Sprintf("%s:%d", db.config.Hostname, db.config.Port),
		DBName:    db.config.Database,
		ParseTime: true,
	}

	var err error
	db.db, err = gorm.Open(mysql.Open(conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("mysql connect failed: %s", err.Error())
	}

	return nil
}

func (db *MySQL) Get() *gorm.DB {
	return db.db
}

func (db *MySQL) AutoMigrate() {
	db.db.AutoMigrate(&schemas.User{}, &schemas.Session{}, &schemas.SteamLink{}, &schemas.CharacterStats{}, &schemas.CharacterStatsFloat{}, &schemas.CharacterStatsInteger{}, &schemas.CharacterStatsString{})
}

func (db *MySQL) PopulateIfFresh() {
	var stat schemas.CharacterStats
	result := db.db.First(&stat)
	if result.RowsAffected > 0 {
		return
	}

	{
		name := schemas.CharacterStats{
			CanBeEmpty:  false,
			Name:        "character_name",
			Description: "Character's name",
			ValueType:   "STR",
			String: schemas.CharacterStatsString{
				Min: 4,
				Max: 32,
			},
		}

		result = db.db.Create(&name)
		if result.Error != nil {
			log.Errorf("DB::MySQL: Failed to create stat: %s", result.Error.Error())
			return
		}
		log.Infof("DB::MySQL: Character name stat created")
	}

	{
		level := schemas.CharacterStats{
			CanBeEmpty:  false,
			Name:        "lvl",
			Description: "Character's level",
			ValueType:   "INT",
			Integer: schemas.CharacterStatsInteger{
				Min:     0,
				Max:     99,
				Default: 0,
			},
		}

		result = db.db.Create(&level)
		if result.Error != nil {
			log.Errorf("DB::MySQL: Failed to create LVL stat: %s", result.Error.Error())
			return
		}
		log.Infof("DB::MySQL: Level stat created")
	}

	{
		stats := []schemas.CharacterStats{
			{
				CanBeEmpty:  false,
				Name:        "str",
				Description: "Strength",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
			{
				CanBeEmpty:  false,
				Name:        "agi",
				Description: "Agility",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
			{
				CanBeEmpty:  false,
				Name:        "vit",
				Description: "Vitality",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
			{
				CanBeEmpty:  false,
				Name:        "int",
				Description: "Intelligence",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
			{
				CanBeEmpty:  false,
				Name:        "dex",
				Description: "Dexterity",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
			{
				CanBeEmpty:  false,
				Name:        "luk",
				Description: "Luck",
				ValueType:   "INT",
				Integer: schemas.CharacterStatsInteger{
					Min:     0,
					Max:     99,
					Default: 0,
				},
			},
		}

		result = db.db.Create(&stats)
		if result.Error != nil {
			log.Errorf("DB::MySQL: Failed to create stats: %s", result.Error.Error())
			return
		}
		log.Infof("DB::MySQL: Default stats created")
	}
}
