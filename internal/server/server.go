package server

import (
	"net/http"
	"time"
)

type Opts struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(opts Opts) *http.Server {
	return &http.Server{
		Addr:         opts.Addr,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
	}
}
