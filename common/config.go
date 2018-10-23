package common

import (
	"time"

	"github.com/BurntSushi/toml"
)

type Configs struct {
	Listen     string
	DirectMail *DirectMailConfig
	Monitor    *MonitorConfig
	HTTP       *httpConfig
}

type DirectMailConfig struct {
	Format           string
	Version          string
	AccessKeyId      string
	AccessKeySecret  string
	SignatureMethod  string
	SignatureVersion string
	RegionId         string
	AccountName      string
	AddressType      string
	ReplyToAddress   string
	FromAlias        string
	ClickTrace       string
	Timeout          Duration
}

type MonitorConfig struct {
	Namespace string
	Subsystem string
}

type httpConfig struct {
	Listen string
}

// Config 全局配置信息
var Config *Configs

// InitConfig 加载配置
func InitConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	Config = config
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}
	return config, nil
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) (err error) {
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

func (d *Duration) D() time.Duration {
	return d.Duration
}
