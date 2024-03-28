// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"fmt"
	"strings"
	"testing"
)

func joinStrings(strs []string) string {
	return strings.Join(strs, "|")
}

// go test -v -cover -count=1 -test.cpu=1 -run=^TestNewSendEmailRequest$
func TestNewSendEmailRequest(t *testing.T) {
	email := &Email{
		TemplateID:    100,
		To:            []string{"to"},
		Cc:            []string{"cc"},
		Bcc:           []string{"bcc"},
		SubjectParams: map[string]string{},
		ContentParams: map[string]string{},
	}

	options := []SendOption{}
	request, _ := newSendEmailRequest(email, options)
	if request.Email.TemplateId != email.TemplateID {
		t.Fatalf("request.Email.TemplateId %d != email.TemplateID %d", request.Email.TemplateId, email.TemplateID)
	}

	if joinStrings(request.Email.To) != joinStrings(email.To) {
		t.Fatalf("request.Email.To %+v != email.To %+v", request.Email.To, email.To)
	}

	if joinStrings(request.Email.Cc) != joinStrings(email.Cc) {
		t.Fatalf("request.Email.Cc %+v != email.Cc %+v", request.Email.Cc, email.Cc)
	}

	if joinStrings(request.Email.Bcc) != joinStrings(email.Bcc) {
		t.Fatalf("request.Email.Bcc %+v != email.Bcc %+v", request.Email.Bcc, email.Bcc)
	}

	if fmt.Sprintf("%p", request.Email.SubjectParams) != fmt.Sprintf("%p", email.SubjectParams) {
		t.Fatalf("request.Email.SubjectParams %+v != email.SubjectParams %+v", request.Email.SubjectParams, email.SubjectParams)
	}

	if fmt.Sprintf("%p", request.Email.ContentParams) != fmt.Sprintf("%p", email.ContentParams) {
		t.Fatalf("request.Email.ContentParams %+v != email.ContentParams %+v", request.Email.ContentParams, email.ContentParams)
	}
}
