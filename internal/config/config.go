package config

import (
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
)

type Config struct {
    DatabaseURL string
    Port string
}

func LoadConfig() *Config {
    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    return &Config{
        DatabaseURL: viper.GetString("DATABASE_URL"),
        Port: viper.GetString("PORT"),
    }
}
