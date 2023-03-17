package cfg

import "github.com/BurntSushi/toml"

var CFG *Config

type Config struct {
	Server struct {
		BindAddress string `toml:"BindAddress"`
		CliToken    string `toml:"CliToken"`
	} `toml:"Server"`

	MicroServices map[string]struct {
		BindAddr string `toml:"BindAddr"`
		Token    string `toml:"Token"`
	} `toml:"MicroServices"`

	Database struct {
		URL   string `toml:"URL"`
		Limit int    `toml:"Limit"`
	} `toml:"Database"`
	Log struct {
		LogLevel  string `toml:"LogLevel"`
		LogOutput string `toml:"LogOutput"`
		LogPath   string `toml:"LogPath"`
	} `toml:"Log"`
}

func NewConfig(cfgPath string) (*Config, error) {
	cfg := &Config{}
	_, err := toml.DecodeFile(cfgPath, cfg)
	if err != nil {
		return nil, err
	}
	CFG = cfg
	return cfg, nil
}
