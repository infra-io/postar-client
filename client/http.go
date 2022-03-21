// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	postarapi "github.com/avino-plan/api/go-out/postar"
	"google.golang.org/protobuf/proto"
)

// httpClient is a client used http connection.
type httpClient struct {
	client *http.Client
	url    *url.URL
}

// NewHTTPClient return a http client.
// The url may look like: http://127.0.0.1:5897
func NewHTTPClient(client *http.Client, url *url.URL) Client {
	return &httpClient{
		client: client,
		url:    url,
	}
}

// sendEmailURL returns the url of sendEmail api.
func (hc *httpClient) sendEmailURL() string {
	return hc.url.Scheme + "://" + hc.url.Host + "/sendEmail"
}

// sendEmailContentType returns the content type of sendEmail api.
func (hc *httpClient) sendEmailContentType() string {
	return "application/octet-stream"
}

// SendEmail sends an email with given options.
// It returns traceID on success and error on fail.
func (hc *httpClient) SendEmail(ctx context.Context, email *Email, opts ...Option) (string, error) {
	req := &postarapi.SendEmailRequest{
		Email:   toAPIEmail(email),
		Options: toAPIOptions(opts...),
	}

	marshaled, err := proto.Marshal(req)
	if err != nil {
		return "", err
	}

	result, err := hc.client.Post(hc.sendEmailURL(), hc.sendEmailContentType(), bytes.NewReader(marshaled))
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	rsp := new(postarapi.SendEmailResponse)
	err = proto.Unmarshal(body, rsp)
	if err != nil {
		return "", err
	}

	return rsp.TraceId, newError(rsp)
}

// Close closes the grpc client.
func (hc *httpClient) Close() error {
	hc.client.CloseIdleConnections()
	return nil
}
