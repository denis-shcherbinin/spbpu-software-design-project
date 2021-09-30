package api

import "time"

type Opts struct {
	API  APIOpts  `group:"API" namespace:"api" env-namespace:"API"`
	DB   DBOpts   `group:"DB" namespace:"db" env-namespace:"DB"`
	Auth AuthOpts `group:"AUTH" namespace:"auth" env-namespace:"AUTH"`

	Debug bool `long:"debug" env:"DEBUG" required:"true" description:"debug mode"`
}

type APIOpts struct {
	Host         string        `long:"host" env:"HOST" required:"true" description:"host"`
	Port         int           `long:"port" env:"PORT" required:"true" description:"port"`
	ReadTimeout  time.Duration `long:"read-timeout" env:"READ_TIMEOUT" required:"true" description:"api read timeout"`
	WriteTimeout time.Duration `long:"write-timeout" env:"WRITE_TIMEOUT" required:"true" description:"api write timeout"`
}

type DBOpts struct {
	Host     string `long:"host" env:"HOST" required:"true" description:"db host"`
	User     string `long:"user" env:"USER" required:"true" description:"db user"`
	Password string `long:"password" env:"PASSWORD" required:"true" description:"db password"`
	Name     string `long:"name" env:"NAME" required:"true" description:"db name"`
	Port     int    `long:"port" env:"PORT" required:"true" description:"db port"`
}

type AuthOpts struct {
	PasswordSalt string `long:"password-salt" env:"PASSWORD_SALT" required:"true" description:"required password salt"`
}
