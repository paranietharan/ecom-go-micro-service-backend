package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Warning: .env file not found, using default environment variables")
		} else {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}
	return nil
}

func GetEnvStr(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return defaultValue
	}
	return value
}

func GetEnvInt(key string, defaultValue int) int {
	valueStr := GetEnvStr(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Warning: Invalid integer value for %s, using default value %d\n", key, defaultValue)
		return defaultValue
	}
	return value
}
