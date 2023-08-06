package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	AppName string
	Port    int
}

var config AppConfig

func InitConfig(environment string) error {
	// Set the path to the configuration files
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	// Add the environment name as a suffix to the configuration file name
	viper.SetConfigType("yaml")
	viper.SetConfigName(fmt.Sprintf("config_%s", environment))

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the configuration into a struct
	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	config = AppConfig {
		viper.GetString("app.name"),
		viper.GetInt("app.port"),
	}

	return nil
}

func GetConfig() AppConfig {
	return config
}
