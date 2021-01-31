package config

import (
	"github.com/spf13/viper"
	"inspiranesia/system/logging"
)

func ProvideViper(log logging.NougatLoggingProvider) *DefaultConfig {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	defaultConfig := DefaultConfig{}
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		log.Fatalf("Error Unmarshall properties, %v", err)
	}
	log.Info("Configuration loaded !")
	return &defaultConfig
}