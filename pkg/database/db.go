package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"learning-gin/pkg/config"
)

func ConnectDB() *sqlx.DB {
	connection := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v sslmode=disable",
		viper.GetString(config.DBUser),
		viper.GetString(config.DBPassword),
		viper.GetString(config.DBHost),
		viper.GetInt(config.DBPort),
		viper.GetString(config.DBName),
	)

	db, err := sqlx.Connect("postgres", connection)
	if err != nil {
		panic(fmt.Sprintf("Error while connecting to database: %v", err.Error()))
	}

	return db
}
