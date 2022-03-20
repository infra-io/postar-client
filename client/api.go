// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"github.com/FishGoddess/errors"
	baseapi "github.com/avino-plan/api/go-out/base"
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

func newError(rsp *postarapi.SendEmailResponse) error {
	if rsp == nil {
		return errors.New("rsp is nil")
	}

	if rsp.Code == baseapi.ServerCode_OK {
		return nil
	}

	if rsp.Code == baseapi.ServerCode_BAD_REQUEST {
		return errors.BadRequest(errors.New(rsp.Msg))
	}

	if rsp.Code == baseapi.ServerCode_TIMEOUT {
		return errors.Timeout(errors.New(rsp.Msg))
	}

	return errors.Wrap(errors.New(rsp.Msg), int32(rsp.Code))
}

// IsBadRequest returns if err is bad request.
func IsBadRequest(err error) bool {
	return errors.IsBadRequest(err)
}

// IsTimeout returns if err is timeout.
func IsTimeout(err error) bool {
	return errors.IsTimeout(err)
}
