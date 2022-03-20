// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"context"

	postarapi "github.com/avino-plan/api/go-out/postar"
	"google.golang.org/grpc"
)

// grpcClient is a client used grpc connection.
type grpcClient struct {
	conn *grpc.ClientConn
}

// NewGRPCClient return a grpc client.
func NewGRPCClient(conn *grpc.ClientConn) Client {
	return &grpcClient{conn: conn}
}

// SendEmail sends an email with given options.
// It returns traceID on success and error on failed.
func (gc *grpcClient) SendEmail(ctx context.Context, email *Email, opts ...Option) (string, error) {
	client := postarapi.NewPostarServiceClient(gc.conn)

	req := &postarapi.SendEmailRequest{
		Email:   toAPIEmail(email),
		Options: toAPIOptions(opts...),
	}
	rsp, err := client.SendEmail(ctx, req)
	if err != nil {
		return "", err
	}

	// TODO 错误处理
	return rsp.TraceId, nil
}

// Close closes the grpc client.
func (gc *grpcClient) Close() error {
	return gc.conn.Close()
}
