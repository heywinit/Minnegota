package config

var DefaultConfig = ServerNetworkConfig{}

type ServerNetworkConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"host"`
}
