// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"context"

	"github.com/infro-io/postar-client/option"
	grpcx "github.com/infro-io/postar-client/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type grpcEmailService struct {
	spaceID    int
	spaceToken string

	conn *grpc.ClientConn
}

func NewGrpcEmailService(address string, spaceID int, spaceToken string, opts ...option.Option) (EmailService, error) {
	conf := option.NewConfig()

	for _, opt := range opts {
		opt(conf)
	}

	ctx, cancel := context.WithTimeout(context.Background(), conf.Timeout)
	defer cancel()

	dialOptions, err := newGrpcDialOptions(conf)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.DialContext(ctx, address, dialOptions...)
	if err != nil {
		return nil, err
	}

	service := &grpcEmailService{
		spaceID:    spaceID,
		spaceToken: spaceToken,
		conn:       conn,
	}

	return service, nil
}

func newGrpcDialOptions(conf *option.Config) ([]grpc.DialOption, error) {
	var dialOptions []grpc.DialOption

	if conf.CertFile != "" {
		creds, err := grpcx.NewClientTLSFromCert(conf.CertFile)
		if err != nil {
			return nil, err
		}

		dialOptions = append(dialOptions, grpc.WithTransportCredentials(creds))
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return dialOptions, nil
}

func (ges *grpcEmailService) newSendResult(md metadata.MD) *SendResult {
	result := &SendResult{
		TraceID: grpcx.GetTraceID(md),
	}

	return result
}

func (ges *grpcEmailService) SendEmail(ctx context.Context, email *Email) (*SendResult, error) {
	if email == nil {
		return nil, errNilEmail
	}

	ctx = grpcx.SetSpace(ctx, ges.spaceID, ges.spaceToken)
	emailService := newEmailService(ges.conn)
	request := newSendEmailRequest(email)

	md := metadata.Pairs()
	callOptions := []grpc.CallOption{grpc.Header(&md)}

	_, err := emailService.SendEmail(ctx, request, callOptions...)
	if err != nil {
		return nil, err
	}

	return ges.newSendResult(md), nil
}

func (ges *grpcEmailService) Close() error {
	return ges.conn.Close()
}
