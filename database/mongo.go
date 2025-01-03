package database

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Hostname         string `yaml:"hostname"`
	Port             uint32 `yaml:"port"`
	Database         string `yaml:"database"`
	Retry            bool   `yaml:"retry"`
	RetryAttempts    int    `yaml:"attempts"`
	ReconnectTimeout int    `yaml:"reconnect_timeout"`
}

type CreateCollection func() error

type MongoDB struct {
	Config      *MongoDBConfig
	Client      *mongo.Client
	Database    *mongo.Database
	collections map[string]CreateCollection
}

func (d *MongoDB) Init(config *MongoDBConfig) error {
	log.Traceln("MongoDB::Init")
	if config == nil {
		return fmt.Errorf("nil config")
	}
	d.Config = config

	if d.Config.Hostname == "" {
		return fmt.Errorf("mongo hostname is missing")
	}
	if d.Config.Port == 0 {
		return fmt.Errorf("mongo port is missing")
	}

	d.collections = make(map[string]CreateCollection)
	//d.collections["service_requests"] = d.CreateServiceRequestsCollection
	//d.collections["auth_requests"] = d.CreateAuthRequestsCollection
	//d.collections["sessions"] = d.CreateSessionsCollection

	return nil
}

func (d *MongoDB) Connect() error {
	if d.Config == nil {
		return fmt.Errorf("mongo: nil config")
	}
	log.Infof("Mongo: Connecting to %s:%d", d.Config.Hostname, d.Config.Port)

	var uri string
	if d.Config.Username != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", d.Config.Username, d.Config.Password, d.Config.Hostname, d.Config.Port, d.Config.Database)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%d/%s", d.Config.Hostname, d.Config.Port, d.Config.Database)
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	var err error
	d.Client, err = mongo.Connect(context.TODO(), opts)

	if err != nil {
		return fmt.Errorf("Mongo: Failed to connect: %s", err.Error())
	}
	/*
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				return fmt.Er
			}
		}()
	*/

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := d.Client.Database(d.Config.Database).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return fmt.Errorf("mongo: Ping failed: %s", err.Error())
	}
	log.Infoln("Mongo: Connection established")
	d.Database = d.Client.Database(d.Config.Database)

	return nil
}

func (d *MongoDB) Disconnect() error {
	log.Infof("Mongo: Disconnecting")
	if d.Client == nil {
		return fmt.Errorf("nil client")
	}
	if err := d.Client.Disconnect(context.TODO()); err != nil {
		return fmt.Errorf("Mongo: Failed to disconnect properly: %s", err.Error())
	}
	return nil
}
