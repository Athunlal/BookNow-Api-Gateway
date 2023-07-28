package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	Authsvcurl    string `mapstructure:"AUTHSVCURL"`
	Usersvcurl    string `mapstructure:"USERSVCURL"`
	Productsvcurl string `mapstructure:"productsvcurl"`
	Adminsvcurl   string `mapstructure:"adminsvcurl"`
	Trainsvurl    string `mapstructure:"trainsvcurl"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	err = viper.Unmarshal(&cfg)

	return cfg, err

}
