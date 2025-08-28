package config

import (
	"GoIAChat/pkg/utils"
)

func LoadEnv() error {
    return utils.LoadEnvFile()
}

func GetAPIKey() string {
	return utils.GetEnvVariable("API_KEY")
}