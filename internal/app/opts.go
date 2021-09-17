package app

import "time"

type Opts struct {
	API APIOpts `group:"API" namespace:"api" env-namespace:"API"`

	Debug bool `long:"debug" env:"DEBUG" required:"true" description:"debug mode"`
}

type APIOpts struct {
	Host         string        `long:"host" env:"HOST" required:"true" description:"host"`
	Port         int           `long:"port" env:"PORT" required:"true" description:"port"`
	ReadTimeout  time.Duration `long:"read-timeout" env:"READ_TIMEOUT" required:"true" description:"api read timeout"`
	WriteTimeout time.Duration `long:"write-timeout" env:"WRITE_TIMEOUT" required:"true" description:"api write timeout"`
}
