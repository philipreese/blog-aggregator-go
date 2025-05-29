package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl 			string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig-go.json"

func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if (err != nil) {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil;
}

func (config *Config) SetUser(userName string) error {
	config.CurrentUserName = userName;
	val, err := json.Marshal(&config)
	if err != nil {
		return err;
	}

	configFilePath, err := getConfigFilePath()
	if (err != nil) {
		return err
	}

	if err := os.WriteFile(configFilePath, val, 0644); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/" + configFileName, nil;
}