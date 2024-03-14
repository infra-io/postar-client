// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postar

import "testing"

// go test -v -cover -count=1 -test.cpu=1 -run=^TestSendResultTraceID$
func TestSendResultTraceID(t *testing.T) {
	result := &SendResult{traceID: "xxx"}
	if result.TraceID() != result.traceID {
		t.Fatalf("result.TraceID() %s != result.traceID %s", result.TraceID(), result.traceID)
	}
}
