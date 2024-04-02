// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package option

import (
	"testing"
	"time"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithTimeout$
func TestWithTimeout(t *testing.T) {
	conf := &Config{Timeout: 0}
	WithTimeout(time.Second)(conf)

	if conf.Timeout != time.Second {
		t.Fatal("conf.Timeout is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithCertFile$
func TestWithCertFile(t *testing.T) {
	conf := &Config{CertFile: ""}
	WithCertFile("xxx")(conf)

	if conf.CertFile != "xxx" {
		t.Fatal("conf.CertFile is wrong")
	}
}
