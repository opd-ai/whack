// Package config provides configuration loading using Viper.
package config

import (
	"github.com/spf13/viper"
)

// Config holds all game configuration.
type Config struct {
	Game   GameConfig   `mapstructure:"game"`
	Server ServerConfig `mapstructure:"server"`
	Client ClientConfig `mapstructure:"client"`
	Match  MatchConfig  `mapstructure:"match"`
}

// GameConfig holds core game settings.
type GameConfig struct {
	Title  string `mapstructure:"title"`
	Width  int    `mapstructure:"width"`
	Height int    `mapstructure:"height"`
	TPS    int    `mapstructure:"tps"`
	Seed   int64  `mapstructure:"seed"`
	Genre  string `mapstructure:"genre"`
}

// ServerConfig holds server settings.
type ServerConfig struct {
	Address    string `mapstructure:"address"`
	MaxPlayers int    `mapstructure:"max_players"`
	TickRate   int    `mapstructure:"tick_rate"`
}

// ClientConfig holds client settings.
type ClientConfig struct {
	ServerAddress string `mapstructure:"server_address"`
}

// MatchConfig holds match settings.
type MatchConfig struct {
	Stocks    int    `mapstructure:"stocks"`
	TimeLimit int    `mapstructure:"time_limit"`
	Mode      string `mapstructure:"mode"`
}

// Load reads configuration from config.yaml and returns a Config.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("game.title", "Whack")
	viper.SetDefault("game.width", 800)
	viper.SetDefault("game.height", 600)
	viper.SetDefault("game.tps", 60)
	viper.SetDefault("game.seed", 0)
	viper.SetDefault("game.genre", "fantasy")
	viper.SetDefault("server.address", ":7000")
	viper.SetDefault("server.max_players", 4)
	viper.SetDefault("server.tick_rate", 60)
	viper.SetDefault("client.server_address", "localhost:7000")
	viper.SetDefault("match.stocks", 3)
	viper.SetDefault("match.time_limit", 300)
	viper.SetDefault("match.mode", "ffa")
}
