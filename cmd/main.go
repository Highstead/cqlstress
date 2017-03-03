package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	pwd := os.Args[1:]
	config := ReadConfig(pwd)
}

type Config struct {
	Hosts    []string
	Keyspace string
}

func ReadConfig(pwd string) Config {
	var configFile = pwd + "/config.toml"
	_, err := os.Stat(configFile)
	if err != nil {
		log.Fatal("Config file is missing: ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
