// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/avino-plan/postar-client/client"
)

func main() {
	address, err := url.Parse("http://127.0.0.1:5897")
	if err != nil {
		panic(err)
	}

	httpClient := client.NewHttpClient(&http.Client{}, address)
	defer httpClient.Close()

	email := &client.Email{
		Subject:   "测试邮件",
		Receivers: []string{os.Getenv("POSTAR_RECEIVER")},
		BodyType:  "text/html",
		Body:      "<p>邮件内容</p>",
	}
	traceID, err := httpClient.SendEmail(context.Background(), email, client.WithSync(), client.WithTimeout(30*time.Second))
	if err != nil {
		panic(err)
	}

	fmt.Println("TraceID:", traceID)
}
