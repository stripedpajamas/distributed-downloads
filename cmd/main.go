package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/stripedpajamas/distributed-downloads/node"
)

var configFilePath string

func init() {
	flag.StringVar(&configFilePath, "config", "default_config.json", "Path to node configuration file")
	flag.Parse()
}

func main() {
	config := LoadConfiguration(configFilePath)
	fmt.Printf("the config: %v", config)
}

// LoadConfiguration loads config file and parses it
func LoadConfiguration(file string) node.Configuration {
	var config node.Configuration
	configFile, err := os.Open(file)
	if err != nil {
		errors.Wrap(err, "invalid configuration file path")
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
