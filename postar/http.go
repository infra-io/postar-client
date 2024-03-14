// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type httpClient struct {
	address    string
	spaceID    int
	spaceToken string

	client *http.Client
}

func NewHttpClient(address string, spaceID int, spaceToken string, opts ...Option) Client {
	conf := newConfig()

	for _, opt := range opts {
		opt(conf)
	}

	client := &http.Client{
		Transport:     conf.transport,
		CheckRedirect: conf.checkRedirect,
		Jar:           conf.cookieJar,
		Timeout:       conf.timeout,
	}

	httpClient := &httpClient{
		address:    address,
		client:     client,
		spaceID:    spaceID,
		spaceToken: spaceToken,
	}

	return httpClient
}

func (hc *httpClient) newSendURL() string {
	return "http://" + hc.address + "/api/postar/v1/emails/send"
}

func (hc *httpClient) newSendRequestBody(email *Email, opts []SendOption) (io.Reader, error) {
	options := newSendOptions()

	for _, opt := range opts {
		opt(options)
	}

	body := map[string]any{
		"email":   email,
		"options": options,
	}

	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(bs), nil
}

func (hc *httpClient) newSendRequest(ctx context.Context, email *Email, opts []SendOption) (*http.Request, error) {
	url := hc.newSendURL()

	body, err := hc.newSendRequestBody(email, opts)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Postar-Space-Id", strconv.Itoa(hc.spaceID))
	request.Header.Set("X-Postar-Space-Token", hc.spaceToken)

	return request, nil
}

func (hc *httpClient) newSendResponse(response *http.Response) *SendResult {
	traceID := response.Header.Get("X-Postar-Trace-Id")

	result := &SendResult{
		traceID: traceID,
	}

	return result
}

func (hc *httpClient) SendEmail(ctx context.Context, email *Email, opts ...SendOption) (*SendResult, error) {
	if email == nil {
		return nil, errNilEmail
	}

	request, err := hc.newSendRequest(ctx, email, opts)
	if err != nil {
		return nil, err
	}

	response, err := hc.client.Do(request)
	if err != nil {
		return nil, err
	}

	return hc.newSendResponse(response), nil
}
