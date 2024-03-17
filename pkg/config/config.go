package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Информация о полях берется из .env файла
const (
	DBUser     = "DBUser"
	DBPassword = "DBPassword"
	DBHost     = "DBHost"
	DBPort     = "DBPort"
	DBName     = "DBName"
)

func InitConfig() {
	path, _ := os.Getwd()

	path = filepath.Join(path, "..")
	path = filepath.Join(path, "Cycle-Backend1")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Config initialization error: %v", err.Error()))
	}
}
