// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"testing"
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithTimeout$
func TestWithTimeout(t *testing.T) {
	conf := &config{timeout: 0}
	WithTimeout(time.Second)(conf)

	if conf.timeout != time.Second {
		t.Fatal("conf.timeout is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithCertFile$
func TestWithCertFile(t *testing.T) {
	conf := &config{certFile: ""}
	WithCertFile("xxx")(conf)

	if conf.certFile != "xxx" {
		t.Fatal("conf.certFile is wrong")
	}
}
