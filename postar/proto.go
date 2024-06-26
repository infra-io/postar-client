// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	postarapi "github.com/infra-io/postar/api/genproto/postar/v1"
	"google.golang.org/grpc"
)

func newSendEmailRequest(email *Email) *postarapi.SendEmailRequest {
	request := &postarapi.SendEmailRequest{
		Email: &postarapi.Email{
			TemplateId:    email.TemplateID,
			To:            email.To,
			Cc:            email.Cc,
			Bcc:           email.Bcc,
			SubjectParams: email.SubjectParams,
			ContentParams: email.ContentParams,
		},
	}

	return request
}

func newEmailService(conn grpc.ClientConnInterface) postarapi.EmailServiceClient {
	return postarapi.NewEmailServiceClient(conn)
}
