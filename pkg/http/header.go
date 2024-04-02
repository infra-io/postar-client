// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package http

import (
	"net/http"
	"strconv"
)

const (
	headerKeyContentType = "Content-Type"
	headerKeySpaceID     = "X-Postar-Space-Id"
	headerKeySpaceToken  = "X-Postar-Space-Token"
	headerKeyTraceID     = "X-Postar-Trace-Id"
)

func SetContentTypeJson(header http.Header) http.Header {
	header.Set(headerKeyContentType, "application/json")
	return header
}

func SetSpace(header http.Header, spaceID int, spaceToken string) http.Header {
	header.Set(headerKeySpaceID, strconv.Itoa(spaceID))
	header.Set(headerKeySpaceToken, spaceToken)
	return header
}

func GetTraceID(header http.Header) string {
	return header.Get(headerKeyTraceID)
}
