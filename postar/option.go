// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"net/http"
	"time"
)

type config struct {
	timeout time.Duration

	transport     http.RoundTripper
	checkRedirect func(req *http.Request, via []*http.Request) error
	cookieJar     http.CookieJar
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

func WithHttpTransport(transport http.RoundTripper) Option {
	return func(conf *config) {
		conf.transport = transport
	}
}

func WithHttpCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) Option {
	return func(conf *config) {
		conf.checkRedirect = checkRedirect
	}
}

func WithHttpCookieJar(jar http.CookieJar) Option {
	return func(conf *config) {
		conf.cookieJar = jar
	}
}

type SendOptions struct {
	Async bool `json:"async"`
}

func newSendOptions() *SendOptions {
	return &SendOptions{
		Async: false,
	}
}

type SendOption func(opts *SendOptions)

func WithSendAsync() SendOption {
	return func(opts *SendOptions) {
		opts.Async = true
	}
}
