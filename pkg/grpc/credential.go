// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"github.com/infro-io/postar-client/pkg/tls"
	"google.golang.org/grpc/credentials"
)

func certServerName(certFile string) (string, error) {
	cert, err := tls.LoadCertificate(certFile)
	if err != nil {
		return "", err
	}

	if len(cert.DNSNames) <= 0 {
		return "", nil
	}

	return cert.DNSNames[0], nil
}

func NewClientTLSFromCert(certFile string) (credentials.TransportCredentials, error) {
	serverName, err := certServerName(certFile)
	if err != nil {
		return nil, err
	}

	return credentials.NewClientTLSFromFile(certFile, serverName)
}
