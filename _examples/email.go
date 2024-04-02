// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/infro-io/postar-client/postar"
)

func main() {
	// You can use NewGrpcEmailService to create a grpc client for sending emails.
	// Change it to NewHttpEmailService so you can create a http client, but notice the ip:port!
	//emailService, err := postar.NewHttpEmailService("127.0.0.1:6897", 100, "")
	emailService, err := postar.NewGrpcEmailService("127.0.0.1:5897", 100, "")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	email := &postar.Email{
		TemplateID: 1000001,
		To:         []string{"xxx@qq.com"},
		SubjectParams: map[string]string{
			"num":     "123",
			"subject": "测试客户端",
		},
		ContentParams: map[string]string{
			"p":   "grpc",
			"img": "-",
		},
	}

	result, err := emailService.SendEmail(ctx, email)
	if err != nil {
		panic(err)
	}

	fmt.Println("trace id:", result.TraceID)
}
