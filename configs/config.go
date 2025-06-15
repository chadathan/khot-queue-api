package configs

import (
	"encoding/json"
	"khot-queue-api/internal/models"
	"log"
	"os"
)

func LoadJsonConfig() models.JsonConfig {
	file, err := os.Open("configs/environment/environment.json")
	if err != nil {
		log.Fatalf("Failed to open config: %v", err)
	}
	defer file.Close()

	var cfg models.JsonConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
	return cfg
}
