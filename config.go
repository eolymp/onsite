package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Network string       `yaml:"network"`
	Listen  string       `yaml:"listen"`
	Rules   []ConfigRule `yaml:"rules"`
}

type ConfigRule struct {
	Allow string `yaml:"allow"`
	Deny  string `yaml:"deny"`
	Ports []int  `yaml:"ports"`
}

func ParseConfig(filename string) (*Config, error) {
	c := &Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}
