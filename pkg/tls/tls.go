// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tls

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

func LoadCertificate(certFile string) (*x509.Certificate, error) {
	certFileBytes, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	certPem, _ := pem.Decode(certFileBytes)
	return x509.ParseCertificate(certPem.Bytes)
}

func NewCertPool(certFile string) (*x509.CertPool, error) {
	certFileBytes, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(certFileBytes)

	return pool, nil
}
