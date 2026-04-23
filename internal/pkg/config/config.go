package config

import (
	"io"
	"gopkg.in/yaml.v3"
)

type Config struct {
	P struct {
		Type string 
	} 
}

func Parse(r io.Reader) (Config, error) {
	var cfg Config
	dec := yaml.NewDecoder(r)
	err := dec.Decode(&cfg)
	return cfg, err
}
