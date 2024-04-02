// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package option

import (
	"time"
)

type Config struct {
	Timeout  time.Duration
	CertFile string
}

func NewConfig() *Config {
	return &Config{
		Timeout: 10 * time.Second,
	}
}

type Option func(conf *Config)

func WithTimeout(timeout time.Duration) Option {
	return func(conf *Config) {
		conf.Timeout = timeout
	}
}

func WithCertFile(certFile string) Option {
	return func(conf *Config) {
		conf.CertFile = certFile
	}
}
