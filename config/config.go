package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/sepehrmohseni/go-web-boilerplate/utils"
	"github.com/spf13/viper"
)

type Database struct {
	Address  string
	Port     uint
	Username string
	Database string
	Secret   string
	Driver   string
}

type App struct {
	Version string
	Port    string
	BaseURL string
}

type Config struct {
	Database     Database
	App          App
}

var AppConfig *Config

func GetConfig(path string) (*Config, error) {
	fileName := filepath.Base(path)
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(path)
	splicedFileName := strings.Split(fileName, `.`)
	conf := viper.New()
	conf.SetConfigName(splicedFileName[0])
	conf.SetConfigType(splicedFileName[1])
	conf.AddConfigPath(dir)
	if err := conf.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("cant read config file: %v", err)
	}
	if AppConfig == nil {
		utils.Mutex.Lock()
		AppConfig = &Config{}
		if err := conf.Unmarshal(AppConfig); err != nil {
			return nil, fmt.Errorf("cant create config struct: %v", err)
		}
		utils.Mutex.Unlock()
	}
	return AppConfig, nil
}

func (c *Config) GetDatabaseURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Secret,
		c.Database.Address,
		c.Database.Port,
		c.Database.Database)
}
