package config

import (
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
)

// Represents server configurations
//
// Fields:
// - DatabaseURL (string): url to connect to DB
// - Port (string): Specified port for server to listen
//
// Since 0.0.1
type Config struct {
    DatabaseURL string
    Port string
}

// Initialize config from the .env
//
// Returns: *Config The server configuration
//
// Since: 0.0.1
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
