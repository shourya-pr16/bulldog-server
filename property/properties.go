package property

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

var (
	cfg  *Config
	once sync.Once
)

func ReadProps() *Config {
	once.Do(func() {
		cfg = &Config{}
		data, err := os.ReadFile("property/config.yaml")
		if err != nil {
			fmt.Printf("config.yaml could not be read, error : %#v\n", err)
			panic("Config file read error")
		}
		yaml.Unmarshal(data, cfg)
	})
	return cfg
}
