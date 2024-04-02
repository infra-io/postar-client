// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"context"
	"errors"
)

var (
	errNilEmail = errors.New("postar: email is nil")
)

type Email struct {
	TemplateID    int64             `json:"template_id"`
	To            []string          `json:"to"`
	Cc            []string          `json:"cc"`
	Bcc           []string          `json:"bcc"`
	SubjectParams map[string]string `json:"subject_params"`
	ContentParams map[string]string `json:"content_params"`
}

type SendResult struct {
	TraceID string
}

type EmailService interface {
	SendEmail(ctx context.Context, email *Email) (*SendResult, error)
	Close() error
}
