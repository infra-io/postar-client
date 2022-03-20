// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"testing"
	"time"
)

// go test -v -cover -run=^TestWithSync$
func TestWithSync(t *testing.T) {
	async := true
	opts := &Options{
		Async: &async,
	}

	WithSync()(opts)
	if *opts.Async != false {
		t.Errorf("*opts.Async %+v != false", *opts.Async)
	}
}

// go test -v -cover -run=^TestWithAsync$
func TestWithAsync(t *testing.T) {
	async := false
	opts := &Options{
		Async: &async,
	}

	WithAsync()(opts)
	if *opts.Async != true {
		t.Errorf("*opts.Async %+v != false", *opts.Async)
	}
}

// go test -v -cover -run=^TestWithTimeout$
func TestWithTimeout(t *testing.T) {
	timeout := time.Duration(0)
	opts := &Options{
		Timeout: &timeout,
	}

	WithTimeout(3 * time.Second)(opts)
	if *opts.Timeout != 3*time.Second {
		t.Errorf("*opts.Timeout %+v != 3*time.Second", *opts.Timeout)
	}
}
