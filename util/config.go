package util

import "github.com/spf13/viper"

// Stores all configuration vars of the application.
type Config struct {
	FlightApiKey  string `mapstructure:"GOOGLE_FLIGHT_KEY"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DatabaseUri   string `mapstructure:"DATABASE_URI"`
	DatabaseName  string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)

	return
}
