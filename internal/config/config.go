package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

func LoadConfig() *mysql.Config {

	DB := &mysql.Config{
		User:      GetEnv("DB_USER", "root"),
		Passwd:    GetEnv("DB_PASSWORD", ""), // Modify
		Net:       "tcp",
		Addr:      GetEnv("DB_ADDR", "localhost:3306"),
		DBName:    GetEnv("DB_NAME", "bootcamp_db"),
		ParseTime: true,
	}
	return DB
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
