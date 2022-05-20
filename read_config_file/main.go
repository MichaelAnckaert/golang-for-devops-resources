package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_VAR  string
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	Name     string
	User     string
	Password string
}

func main() {
	var configuration Configurations
	
        viper.SetConfigName("config")
        viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/myapp")
        viper.SetConfigType("yml")
	viper.AutomaticEnv()

	// Set default variable values
	viper.SetDefault("database.name", "test_db")


        if err := viper.ReadInConfig(); err != nil {
                fmt.Printf("Error reading config file, %s", err)
        }

        err := viper.Unmarshal(&configuration)
        if err != nil {
                fmt.Printf("Unable to decode into struct, %v", err)
        }

        fmt.Println("Database is\t", configuration.Database.Name)
        fmt.Println("Port is\t\t", configuration.Server.Port)
        fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)
}
