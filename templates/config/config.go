package config

import "github.com/spf13/viper"

type Config struct {
	DbConnectionString            string `mapstructure:"db_connection_string"`
	PublishBrokerConnectionString string `mapstructure:"publish_broker_connection_string"`
	ConsumeBrokerConnectionString string `mapstructure:"consumer_broker_connection_string"`
	ApiPort                       int    `mapstructure:"api_port"`
}

var AppConfig *Config

func LoadConfig() {

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	AppConfig.DbConnectionString = viper.Get("database.db_connection_string").(string)
	AppConfig.PublishBrokerConnectionString = viper.Get("message_brokers.publish_broker_connection_string").(string)
	AppConfig.ConsumeBrokerConnectionString = viper.Get("message_brokers.consume_broker_connection_string").(string)
	AppConfig.ApiPort = viper.Get("api.port").(int)

}
