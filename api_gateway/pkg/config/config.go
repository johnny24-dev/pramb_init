package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	Authsvcurl   string `mapstructure:"AUTHSVCURL"`
	Usersvcurl   string `mapstructure:"USERSVCURL"`
	Adminsvcurl  string `mapstructure:"ADMINSVCURL"`
	Trainsvurl   string `mapstructure:"TRAINSVCURL"`
	Bookingsvurl string `mapstructure:"BOOKINGSVCURL"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	err := viper.Unmarshal(&cfg)
	LoadEnv()

	return cfg, err
}
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Env File")
	}

}
