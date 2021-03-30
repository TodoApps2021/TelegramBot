package kafka

import "github.com/spf13/viper"

type ConfigProducer struct {
	BootstrapServers string
}

func ReadFromYAMLProducer() *ConfigProducer {
	return &ConfigProducer{
		BootstrapServers: viper.GetString("kafka.url"),
	}
}
