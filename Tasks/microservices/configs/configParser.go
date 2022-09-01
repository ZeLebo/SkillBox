package configs

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
}

func RootDir() string {
	ex, err := os.Getwd()
	if err != nil {
		log.Fatalf("Cannot parse the working directory")
	}
	exPath := filepath.Base(ex)
	switch exPath {
	case "microservices":
		return "configs"
	case "cmd":
		return filepath.Join("..", "configs")
	case "app":
		return filepath.Join("..", "..", "configs")
	default:
		fmt.Println(exPath)
		return ""
	}
}

func InitConfig() error {
	// find configs folder
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
