package pkgs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Env struct {
	DBUserName string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
	Port       string
}

var EnvironmentValues Env

func config() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
func LoadEnvironment() Env {
	config()
	EnvironmentValues = Env{
		DBUserName: viper.Get("APP_DB_USERNAME").(string),
		DBPassword: viper.Get("APP_DB_PASSWORD").(string),
		DBName:     viper.Get("APP_DB_NAME").(string),
		DBPort:     viper.Get("APP_DB_PORT").(string),
		DBHost:     viper.Get("APP_DB_HOST").(string),
		Port:       viper.Get("APP_Port").(string),
	}
	return EnvironmentValues
}

func Must(e Env) error {
	fmt.Println(e, viper.Get("APP_DB_USERNAME"))
	if e.DBHost == "" || e.DBPort == "" || e.DBName == "" || e.DBUserName == "" || e.DBPassword == "" || e.Port == "" {
		panic(fmt.Errorf("Missing enivronment variable"))
	}
	return nil
}
