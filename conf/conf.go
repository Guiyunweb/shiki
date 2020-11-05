package conf

import "github.com/spf13/viper"

var Conf *viper.Viper

func Init() error {
	Conf = viper.New()
	Conf.SetConfigFile("conf/app.toml")
	return Conf.ReadInConfig()
}
