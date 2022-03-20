// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import postarapi "github.com/avino-plan/api/go-out/postar"

func toAPIEmail(email *Email) *postarapi.Email {
	if email == nil {
		return nil
	}

	return &postarapi.Email{
		Receivers: email.receivers,
		Subject:   email.subject,
		BodyType:  email.bodyType,
		Body:      email.body,
	}
}
