// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -cover -run=^TestToAPIEmail$
func TestToAPIEmail(t *testing.T) {
	if toAPIEmail(nil) != nil {
		t.Error("toAPIEmail(nil) != nil")
	}

	email := &Email{
		Subject:   "Subject",
		Receivers: []string{"Receiver"},
		BodyType:  "BodyType",
		Body:      "Body",
	}

	apiEmail := toAPIEmail(email)
	if apiEmail.Subject != email.Subject {
		t.Errorf("apiEmail.Subject %s != email.Subject %s", apiEmail.Subject, email.Subject)
	}

	if fmt.Sprintf("%+v", apiEmail.Receivers) != fmt.Sprintf("%+v", email.Receivers) {
		t.Errorf("apiEmail.BodyType %+v != email.BodyType %+v", apiEmail.Receivers, email.Receivers)
	}

	if apiEmail.BodyType != email.BodyType {
		t.Errorf("apiEmail.BodyType %s != email.BodyType %s", apiEmail.BodyType, email.BodyType)
	}

	if apiEmail.Body != email.Body {
		t.Errorf("apiEmail.Body %s != email.Body %s", apiEmail.Body, email.Body)
	}
}

// go test -v -cover -run=^TestToAPIOptions$
func TestToAPIOptions(t *testing.T) {
	opts := toAPIOptions(WithAsync(), WithTimeout(3*time.Second))
	if opts.Async != true {
		t.Errorf("opts.Async %+v != true", opts.Async)
	}

	if int64(opts.TimeoutMillis) != (3 * time.Second).Milliseconds() {
		t.Errorf("int64(opts.TimeoutMillis) %d != (3 * time.Second).Milliseconds()", opts.TimeoutMillis)
	}
}
