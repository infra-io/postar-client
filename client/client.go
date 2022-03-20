// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"context"
	"time"
)

// Email is the struct of email.
type Email struct {
	Subject   string
	Receivers []string
	BodyType  string
	Body      string
}

// Options is the Options of sending one email.
type Options struct {
	Async   *bool          // The mode of sending one email.
	Timeout *time.Duration // The timeout of sending one email.
}

// Option applies an option to opts.
type Option func(opts *Options)

// WithSync uses sync mode to send emails.
func WithSync() Option {
	return func(opts *Options) {
		async := false
		opts.Async = &async
	}
}

// WithAsync uses async mode to send emails.
func WithAsync() Option {
	return func(opts *Options) {
		async := true
		opts.Async = &async
	}
}

// WithTimeout sets timeout of sending emails.
func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = &timeout
	}
}

// Client is an interface of Postar service client.
type Client interface {
	// SendEmail sends an email with given options.
	// It returns traceID on success and error on fail.
	SendEmail(ctx context.Context, email *Email, opts ...Option) (string, error)

	// Close closes the grpc client.
	Close() error
}
