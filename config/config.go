package config

import (
	"os"
	"strconv"
)

type Config struct {
	SqlAddress string
	ServerPort uint16
}

func LoadConfig() Config {
	cfg := Config{
		SqlAddress: "root:@tcp(127.0.0.1:3306)/task_management?charset=utf8&parseTime=True&loc=Local",
		ServerPort: 8080,
	}

	if sqlAddr, exists := os.LookupEnv("SQL_ADDR"); exists {
		cfg.SqlAddress = sqlAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
