// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	postarapi "github.com/avino-plan/api/go-out/postar"
)

func toAPIEmail(email *Email) *postarapi.Email {
	if email == nil {
		return nil
	}

	return &postarapi.Email{
		Receivers: email.Receivers,
		Subject:   email.Subject,
		BodyType:  email.BodyType,
		Body:      email.Body,
	}
}

func toAPIOptions(opts ...Option) *postarapi.SendEmailOptions {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	result := new(postarapi.SendEmailOptions)
	if o.Async != nil {
		result.Async = *o.Async
	}

	if o.Timeout != nil {
		result.TimeoutMillis = int32((*o.Timeout).Milliseconds())
	}
	return result
}
