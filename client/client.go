// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"context"

	"github.com/avino-plan/postar-client/options"
)

// Email is the struct of email.
type Email struct {
	subject   string
	receivers []string
	bodyType  string
	body      string
}

// NewEmail returns an Email instance.
func NewEmail() *Email {
	return new(Email)
}

// Client is an interface of Postar service client.
type Client interface {
	SendEmail(ctx context.Context, email *Email, opts ...options.Option) (string, error)
}
