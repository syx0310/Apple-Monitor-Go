package apple

import (
	"github.com/spf13/viper"
)

type Device struct {
	Name                  string            `mapstructure:"name"`
	ProductID             string            `mapstructure:"product_id"`
	Location              string            `mapstructure:"location"`
	Region                string            `mapstructure:"region"`
	QueryParams           map[string]string `mapstructure:"query_params"`
	Crontab               string            `mapstructure:"crontab"`
	BarkKey               string            `mapstructure:"bark_key"`
	BarkAPIURL            string            `mapstructure:"bark_api_url"`
	WeComURL              string            `mapstructure:"wecom_url"`
	StoreWhitelistKeyword []string          `mapstructure:"store_whitelist_keyword"`
}

type Config struct {
	Devices []Device `mapstructure:"devices"`
}

var AppConfig Config

func InitConfig() error {
	return viper.Unmarshal(&AppConfig)
}
