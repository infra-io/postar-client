// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	metadataKeySpaceID    = "postar.space_id"
	metadataKeySpaceToken = "postar.space_token"
	metadataKeyTraceID    = "postar.trace_id"
)

func SetSpace(md metadata.MD, spaceID int, spaceToken string) metadata.MD {
	md.Set(metadataKeySpaceID, strconv.Itoa(spaceID))
	md.Set(metadataKeySpaceToken, spaceToken)
	return md
}

func getMetadata(md metadata.MD, key string) string {
	values := md.Get(key)
	if len(values) <= 0 {
		return ""
	}

	return values[0]
}

func GetTraceID(md metadata.MD) string {
	return getMetadata(md, metadataKeyTraceID)
}
