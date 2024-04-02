// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"time"
)

type config struct {
	timeout  time.Duration
	certFile string
}

func newConfig() *config {
	return &config{
		timeout: 10 * time.Second,
	}
}

type Option func(conf *config)

func WithTimeout(timeout time.Duration) Option {
	return func(conf *config) {
		conf.timeout = timeout
	}
}

func WithCertFile(certFile string) Option {
	return func(conf *config) {
		conf.certFile = certFile
	}
}
