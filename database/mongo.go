package database

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

type MongoDB struct{}

func (d *MongoDB) Init(config *MongoDBConfig) error {
	return nil
}
