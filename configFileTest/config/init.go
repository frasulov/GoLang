package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Configuration Configurations
func Init() {
	viper.SetConfigName("config-dev.yml")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil{
		fmt.Printf("Error reading config file: %v", err.Error())
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil{
		fmt.Printf("Unable to decode into struct %v", err.Error())
	}

	activeProfile := Configuration.Profile.Active
	if activeProfile != "dev" {
		viper.SetConfigName("config-"+activeProfile+".yml")
		if err := viper.ReadInConfig(); err != nil{
			fmt.Printf("Error reading config file: %v", err.Error())
		}
		err := viper.Unmarshal(&Configuration)
		if err != nil{
			fmt.Printf("Unable to decode into struct %v", err.Error())
		}
	}
}
