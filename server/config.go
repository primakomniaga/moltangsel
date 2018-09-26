package server

var Config struct {
	Server         ServerConfig   `yaml:"server"`
	DatabaseMaster DatabaseConfig `yaml:"databaseMaster"`
	DatabaseSlave  DatabaseConfig `yaml:"databaseSlave"`
	Redis          RedisConfig    `yaml:"redis"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Address  string `yaml:"address"`
	Database string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Address string `yaml:"address"`
}
