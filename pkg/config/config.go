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
		URL string `toml:"URL"`
	} `toml:"Database"`
	Log struct {
		LogLevel  string `toml:"LogLevel"`
		LogOutput string `toml:"LogOutput"`
		LogPath   string `toml:"LogPath"`
	} `toml:"Log"`
	PostAPI struct {
		PostToken string `toml:"PostToken"`
		PostAuth  string `toml:"PostAuth"`
		URL       string `toml:"URL"`
	} `toml:"PostAPI"`
	Addresses struct {
		Quality    map[string]string `toml:"Quality"`
		Validation map[string]string `toml:"Validation"`
		QcGeo      map[string]string `toml:"QC_geo"`
		QcComplete map[string]string `toml:"QC_complete"`
		QcHouse    map[string]string `toml:"QC_house"`
		Qc         map[string]string `toml:"QC"`
		SSCode     map[string]string `toml:"SS_code"`
	} `toml:"Addresses"`
	DaDataAPI struct {
		APIKey    string `toml:"APIKey"`
		SecretKEY string `toml:"SecretKEY"`
		LiveTime  int    `toml:"LiveTime"`
	} `toml:"DaDataAPI"`
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