package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Network          string       `yaml:"network"`
	Listen           string       `yaml:"listen"`
	AllowedDomains   []string     `yaml:"allowed_domains"`
	ForbiddenDomains []string     `yaml:"forbidden_domains"`
	AllowedIP        []string     `yaml:"allowed_ip"`
	ForbiddenIP      []string     `yaml:"forbidden_ip"`
	AllowedPorts     []int        `yaml:"allowed_ports"`
	ForbiddenPorts   []int        `yaml:"forbidden_ports"`
	AllowResolvedIPs bool         `yaml:"allow_resolved_ips"`
	Rules            []ConfigRule `yaml:"rules"`
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
