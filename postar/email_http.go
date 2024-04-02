// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	httpx "github.com/infro-io/postar-client/pkg/http"
	tlsx "github.com/infro-io/postar-client/pkg/tls"
)

type httpEmailService struct {
	address    string
	spaceID    int
	spaceToken string

	client *http.Client
}

func NewHttpEmailService(address string, spaceID int, spaceToken string, opts ...Option) (EmailService, error) {
	conf := newConfig()

	for _, opt := range opts {
		opt(conf)
	}

	transport, err := newHttpTransport(conf)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   conf.timeout,
	}

	service := &httpEmailService{
		address:    address,
		spaceID:    spaceID,
		spaceToken: spaceToken,
		client:     client,
	}

	return service, nil
}

func newHttpTransport(conf *config) (*http.Transport, error) {
	transport := new(http.Transport)

	if conf.certFile != "" {
		certPool, err := tlsx.NewCertPool(conf.certFile)
		if err != nil {
			return nil, err
		}

		transport.TLSClientConfig = &tls.Config{RootCAs: certPool}
	}

	return transport, nil
}

func (hes *httpEmailService) newSendURL() string {
	return "http://" + hes.address + "/api/postar/v1/emails/send"
}

func (hes *httpEmailService) newSendRequestBody(email *Email) (io.Reader, error) {
	request := newSendEmailRequest(email)

	bs, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bs), nil
}

func (hes *httpEmailService) newSendRequest(ctx context.Context, email *Email) (*http.Request, error) {
	url := hes.newSendURL()

	body, err := hes.newSendRequestBody(email)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	httpx.SetContentTypeJson(request.Header)
	httpx.SetSpace(request.Header, hes.spaceID, hes.spaceToken)

	return request, nil
}

func (hes *httpEmailService) newSendResult(response *http.Response) *SendResult {
	result := &SendResult{
		TraceID: httpx.GetTraceID(response.Header),
	}

	return result
}

func (hes *httpEmailService) SendEmail(ctx context.Context, email *Email) (*SendResult, error) {
	if email == nil {
		return nil, errNilEmail
	}

	request, err := hes.newSendRequest(ctx, email)
	if err != nil {
		return nil, err
	}

	response, err := hes.client.Do(request)
	if err != nil {
		return nil, err
	}

	return hes.newSendResult(response), nil
}

func (hes *httpEmailService) Close() error {
	return nil
}
