// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	grpcx "github.com/infra-io/servicex/net/grpc"
)

type grpcClient struct {
	spaceID    int
	spaceToken string

	conn *grpc.ClientConn
}

func NewGrpcClient(address string, spaceID int, spaceToken string, opts ...Option) (Client, error) {
	conf := newConfig()

	for _, opt := range opts {
		opt(conf)
	}

	ctx, cancel := context.WithTimeout(context.Background(), conf.timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, conf.grpcDialOptions...)
	if err != nil {
		return nil, err
	}

	client := &grpcClient{
		spaceID:    spaceID,
		spaceToken: spaceToken,
		conn:       conn,
	}

	return client, nil
}

func (gc *grpcClient) newSendResult(md metadata.MD) *SendResult {
	traceID := grpcx.GetMetadata(md, grpcx.ServiceKeyTraceID)

	result := &SendResult{
		traceID: traceID,
	}

	return result
}

func (gc *grpcClient) SendEmail(ctx context.Context, email *Email, opts ...SendOption) (*SendResult, error) {
	if email == nil {
		return nil, errNilEmail
	}

	emailService := newEmailService(gc.conn)
	request, sendOptions := newSendEmailRequest(email, opts)

	var md metadata.MD
	callOptions := append(sendOptions.GrpcCallOptions, grpc.Header(&md))

	_, err := emailService.SendEmail(ctx, request, callOptions...)
	if err != nil {
		return nil, err
	}

	return gc.newSendResult(md), nil
}

func (gc *grpcClient) Close() error {
	return gc.conn.Close()
}
