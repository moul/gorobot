package gorobot

import (
    "fmt"
    "os"
    "json"
    "io/ioutil"
)

type Config struct {
	Logs ConfigLogs
	Module ConfigModule
	Servers map[string] ConfigServer
}

type ConfigLogs struct {
	Enable bool
	Directory string
	RecordEvents bool
	RecordMemoryUsage bool
	RecordStatistics bool
}

type ConfigModule struct {
	Interface string
}

type ConfigServer struct {
	Name string
	Host string
	Nickname string
	Realname string
	Username string
	Password string
	Channels map[string] ConfigChannel
}

type ConfigChannel struct {
	Name string
	Password string
	Master bool
}

// Returns a new configuration from file pointed by path
func NewConfig(path string) (*Config) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("Configuration error: %v\n", e)
		os.Exit(1)
	}
	var config Config
	err := json.Unmarshal(file, &config)
	if err != nil {
		fmt.Printf("Configuration error: %s\n", err)
		os.Exit(1)
	}
	return &config
}
