package config

type LogConfig struct {
	Path  string `yaml:"path"`
	Save  uint   `yaml:"save"`
}