package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFilePath = "./.gatorconfig.json"

type Config struct {
    DbUrl string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
    data, err := os.ReadFile(configFilePath)
    if err != nil {
        log.Fatal(err)
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        log.Fatal(err)
    }
    return config
}

func write(config Config) {
    data, err := json.Marshal(config)
    if err != nil {
        log.Fatal(err)
    }

    os.WriteFile(configFilePath, data, 0660)
}

func (config Config) SetUser(newName string) {
    config.CurrentUserName = newName
    write(config)
}