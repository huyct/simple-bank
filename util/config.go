package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver  string `mapstructure:"dbDriver"`
	DBSource  string `mapstructure:"dbSource"`
	ServerURL string `mapstructure:"serverURL"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
