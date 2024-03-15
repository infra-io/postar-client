// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type config struct {
	timeout time.Duration

	transport     http.RoundTripper
	checkRedirect func(req *http.Request, via []*http.Request) error
	cookieJar     http.CookieJar

	grpcDialOptions []grpc.DialOption
}

func newConfig() *config {
	insecureOpt := grpc.WithTransportCredentials(insecure.NewCredentials())

	return &config{
		timeout:         10 * time.Second,
		grpcDialOptions: []grpc.DialOption{insecureOpt},
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

func WithGrpcDialOptions(grpcOpts ...grpc.DialOption) Option {
	return func(conf *config) {
		conf.grpcDialOptions = append(conf.grpcDialOptions, grpcOpts...)
	}
}

type SendOptions struct {
	Async           bool `json:"async"`
	GrpcCallOptions []grpc.CallOption
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

func WithGrpcCallOptions(grpcOpts ...grpc.CallOption) SendOption {
	return func(opts *SendOptions) {
		opts.GrpcCallOptions = append(opts.GrpcCallOptions, grpcOpts...)
	}
}
