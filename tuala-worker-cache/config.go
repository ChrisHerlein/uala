package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BeanstalkHost string `required:"true" split_words:"true"`
	RedisHost     string `required:"true" split_words:"true"`

	PgHost     string `required:"true" split_words:"true"`
	PgPort     string `required:"true" split_words:"true"`
	PgDb       string `required:"true" split_words:"true"`
	PgUser     string `required:"true" split_words:"true"`
	PgPassword string `required:"true" split_words:"true"`
}

func (c *Config) Init() {
	if err := envconfig.Process("", c); err != nil {
		panic(err)
	}
}
