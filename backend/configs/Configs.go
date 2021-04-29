package configs

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"path/filepath"
)

// We can define more fields for other configs
type Configuration struct {
	DbName     string
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
}

var (
	Config *Configuration
	DBInfo string
)

func InitConfig() {
	if Config == nil {
		Config = NewConfig()
		DBInfo = "host=" + Config.DbHost +" port=" + Config.DbPort + " user=" + Config.DbUser +" dbname=" + Config.DbName + " password=" + Config.DbPassword +" sslmode=disable"
	}
}

func NewConfig() *Configuration {
	var config Configuration
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	configFilePath := dir+"/configs/config.json"
	log.Printf("Config file path: %s", configFilePath)
	err = gonfig.GetConf(configFilePath, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
