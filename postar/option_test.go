// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"google.golang.org/grpc"
)

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithTimeout$
func TestWithTimeout(t *testing.T) {
	conf := &config{timeout: 0}
	WithTimeout(time.Second)(conf)

	if conf.timeout != time.Second {
		t.Fatal("conf.timeout is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithHttpTransport$
func TestWithHttpTransport(t *testing.T) {
	conf := &config{transport: nil}
	WithHttpTransport(http.DefaultTransport)(conf)

	if conf.transport != http.DefaultTransport {
		t.Fatal("conf.transport is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithHttpCheckRedirect$
func TestWithHttpCheckRedirect(t *testing.T) {
	checkRedirect := func(req *http.Request, via []*http.Request) error {
		return nil
	}

	conf := &config{checkRedirect: nil}
	WithHttpCheckRedirect(checkRedirect)(conf)

	if fmt.Sprintf("%p", conf.checkRedirect) != fmt.Sprintf("%p", checkRedirect) {
		t.Fatal("conf.checkRedirect is wrong")
	}
}

type testJar struct{}

func (tj *testJar) SetCookies(u *url.URL, cookies []*http.Cookie) {}

func (tj *testJar) Cookies(u *url.URL) []*http.Cookie {
	return nil
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithHttpCookieJar$
func TestWithHttpCookieJar(t *testing.T) {
	tj := new(testJar)

	conf := &config{cookieJar: nil}
	WithHttpCookieJar(tj)(conf)

	if conf.cookieJar != tj {
		t.Fatal("conf.cookieJar is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithGrpcDialOptions$
func TestWithGrpcDialOptions(t *testing.T) {
	dialOptions := []grpc.DialOption{grpc.WithBlock()}

	conf := &config{grpcDialOptions: nil}
	WithGrpcDialOptions(dialOptions...)(conf)

	if len(conf.grpcDialOptions) != len(dialOptions) {
		t.Fatalf("len(conf.grpcDialOptions) %d != len(dialOptions) %d", len(conf.grpcDialOptions), len(dialOptions))
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithSendAsync$
func TestWithSendAsync(t *testing.T) {
	opts := &SendOptions{Async: false}
	WithSendAsync()(opts)

	if !opts.Async {
		t.Fatal("opts.Async is wrong")
	}
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestWithGrpcCallOptions$
func TestWithGrpcCallOptions(t *testing.T) {
	callOptions := []grpc.CallOption{grpc.EmptyCallOption{}}

	opts := &SendOptions{GrpcCallOptions: nil}
	WithGrpcCallOptions(callOptions...)(opts)

	if len(opts.GrpcCallOptions) != len(callOptions) {
		t.Fatalf("len(opts.GrpcCallOptions) %d != len(callOptions) %d", len(opts.GrpcCallOptions), len(callOptions))
	}
}
