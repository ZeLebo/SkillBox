package domain

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing domain: %s", err.Error())
	}
}

func RootDir() string {
	ex, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot parse the working directory")
	}
	exPath := filepath.Base(ex)
	// navigate to microservices/configs
	switch exPath {
	case "configs":
		return "."
	case "microservices":
		return filepath.Join("configs")
	case "cmd":
		return filepath.Join("..", "configs")
	default:
		return filepath.Join("..", "..", "configs")
	}
}

func InitConfig() error {
	// find domain folder
	viper.AddConfigPath(RootDir())
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func GetIp(ip string) (string, error) {
	res, ok := viper.Get(ip).(string)
	if !ok {
		return "", errors.New("cannot parse data from config")
	}
	return res, nil
}

func GetDatabaseConfig() *DBConfig {
	return &DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}
