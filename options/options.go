// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package options

import "time"

// options is the options of sending one email.
type options struct {
	Async   *bool          // The mode of sending one email.
	Timeout *time.Duration // The timeout of sending one email.
}

type Option func(opts *options)

func WithSync() Option {
	return func(opts *options) {
		async := false
		opts.Async = &async
	}
}

func WithAsync() Option {
	return func(opts *options) {
		async := true
		opts.Async = &async
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *options) {
		opts.Timeout = &timeout
	}
}
