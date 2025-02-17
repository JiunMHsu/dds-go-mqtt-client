package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	HeladeraTopics []string

	PublishTemperature         bool
	PublishTemperatureInterval int
}

var Env = initConfig()

func initConfig() Config {
	godotenv.Load()

	topics := getEnvAsSlice("HELADERA_TOPICS", []string{})
	publishTemp := getEnvAsBool("PUBLISH_TEMPERATURE", false)

	publishTempInterval := getEnvAsInt("PUBLISH_TEMPERATURE_INTERVAL", 5)
	if publishTempInterval < 1 {
		publishTempInterval = 1
	}

	return Config{
		HeladeraTopics:             topics,
		PublishTemperature:         publishTemp,
		PublishTemperatureInterval: publishTempInterval,
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	str := getEnv(key, "")
	value, err := strconv.Atoi(str)
	if err == nil {
		return value
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	str := getEnv(key, "")
	value, err := strconv.ParseBool(str)
	if err == nil {
		return value
	}

	return fallback
}

func getEnvAsSlice(key string, fallback []string) []string {
	str := getEnv(key, "")
	if str == "" {
		return fallback
	}

	str = strings.ReplaceAll(str, "\n", "")
	topics := strings.Split(str, ",")
	for i := range topics {
		topics[i] = strings.TrimSpace(topics[i])
	}

	return topics
}
