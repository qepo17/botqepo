package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TwitterConsumerKey    string `envconfig:"TWITTER_CONSUMER_KEY" default:""`
	TwitterConsumerSecret string `envconfig:"TWITTER_CONSUMER_SECRET" default:""`
	TwitterAccessToken    string `envconfig:"TWITTER_ACCESS_TOKEN" default:""`
	TwitterAccessSecret   string `envconfig:"TWITTER_ACCESS_SECRET" default:""`

	OpenaiBaseUrl  string `envconfig:"OPENAI_BASE_URL" default:"https://api.openai.com"`
	OpenaiToken    string `envconfig:"OPENAI_TOKEN" default:""`
	OpenaiModel    string `envconfig:"OPENAI_MODEL" default:"text-davinci-002"`
	IsOpenAIEnable bool   `envconfig:"IS_OPENAI_ENABLE" default:"true"`

	ZenquotesBaseUrl  string `envconfig:"OPENAI_BASE_URL" default:"https://zenquotes.io"`
	IsZenQuotesEnable bool   `envconfig:"IS_ZENQUOTES_ENABLE" default:"true"`
}

func Init() *Config {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
