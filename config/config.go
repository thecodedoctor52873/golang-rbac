package config

import (
    "log"
    "github.com/spf13/viper"
)

var JWTSecret string
var RedisAddr string
var AuthMethod string

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("json")
    viper.AddConfigPath(".")

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    JWTSecret = viper.GetString("jwt_secret")
    RedisAddr = viper.GetString("redis_addr")
    AuthMethod = viper.GetString("auth_method")
}