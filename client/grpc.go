// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"context"

	postarapi "github.com/avino-plan/api/go-out/postar"
	"github.com/avino-plan/postar-client/options"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	conn *grpc.ClientConn
}

func NewGRPCClient(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{conn: conn}
}

func (gc *GRPCClient) SendEmail(ctx context.Context, email *Email, opts ...options.Option) (string, error) {
	client := postarapi.NewPostarServiceClient(gc.conn)

	req := &postarapi.SendEmailRequest{
		Email:   toAPIEmail(email),
		Options: nil,
	}
	rsp, err := client.SendEmail(ctx, req)
	if err != nil {
		return "", err
	}

	// TODO 错误处理
	return rsp.TraceId, nil
}

func (gc *GRPCClient) Close() error {
	return gc.conn.Close()
}
