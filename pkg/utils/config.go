package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Env          string `json:"env" env-required:"true"`
	PageElements int    `json:"page_elements" env-default:"10"`
	HTTPServer   `json:"http_server"`
	Auth         `json:"auth"`
	Database     `json:"database"`
}

type HTTPServer struct {
	Host string `json:"host" env-default:"localhost"`
	Port string `json:"port" env-default:"8080"`
}

type Auth struct {
	SecretKey string `json:"secret_key"`
	Salt      string `json:"salt"`
}

type Database struct {
	Host     string `json:"host" env-default:"localhost"`
	Port     string `json:"port" env-default:"5432"`
	Username string `json:"username" env-required:"true"`
	Password string `json:"password" env-required:"true"`
	DBName   string `json:"db_name" env-required:"true"`
	SSLMode  string `json:"ssl_mode" env-required:"ture"`
}

func InitConfig(configPath string) {
	// Read file
	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("ошибка чтения файла конфигурации: %v", err)
	}

	// Parse JSON
	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		log.Fatalf("ошибка парсинга JSON: %v", err)
	}

	CONFIG = &config
}
