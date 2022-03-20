// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/avino-plan/postar-client/client"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5897", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcClient := client.NewGRPCClient(conn)
	defer grpcClient.Close()

	email := &client.Email{
		Subject:   "测试邮件",
		Receivers: []string{os.Getenv("POSTAR_RECEIVER")},
		BodyType:  "text/html",
		Body:      "<p>邮件内容</p>",
	}
	traceID, err := grpcClient.SendEmail(context.Background(), email, client.WithAsync(), client.WithTimeout(3*time.Second))
	if err != nil {
		panic(err)
	}

	fmt.Println("TraceID:", traceID)
}
