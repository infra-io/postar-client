// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"context"
	"strconv"

	"google.golang.org/grpc/metadata"
)

const (
	metadataKeySpaceID    = "postar.space_id"
	metadataKeySpaceToken = "postar.space_token"
	metadataKeyTraceID    = "postar.trace_id"
)

func SetSpace(ctx context.Context, spaceID int, spaceToken string) context.Context {
	ctx = metadata.AppendToOutgoingContext(ctx, metadataKeySpaceID, strconv.Itoa(spaceID), metadataKeySpaceToken, spaceToken)
	return ctx
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
