package database

import (
	"fmt"

	_mysql "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Hostname string `yaml:"hostname"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
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
	log.Infof("MySQL: Establishing connection")

	if db.config == nil {
		return fmt.Errorf("mysql: nil config")
	}

	conf := &_mysql.Config{
		User:   db.config.Username,
		Passwd: db.config.Password,
		Addr:   fmt.Sprintf("%s:%d", db.config.Hostname, db.config.Port),
		DBName: db.config.Database,
	}

	var err error
	db.db, err = gorm.Open(mysql.Open(conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}
