package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var configFile = "/home/zhujunheng/GolandProjects/Snake-Go/config/config.json"

func loadConfig(filePath string) (map[string]interface{}, error) {
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var config map[string]interface{}

	if err := json.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfig(key string) (interface{}, error) {
	config, err := loadConfig(configFile)
	if err != nil {
		return nil, err
	}
	value, ok := config[key]
	if !ok {
		return nil, errors.New("config key not found")
	}
	return value, nil
}
