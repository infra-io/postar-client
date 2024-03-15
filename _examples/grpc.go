// Copyright 2022 FishGoddess.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/infro-io/postar-client/postar"
)

func main() {
	client := postar.NewGrpcClient("127.0.0.1:6897", 100, "space_token")

	ctx := context.Background()
	email := &postar.Email{
		TemplateID: 1000000,
		To:         []string{"xxx@abc.com"},
	}

	result, err := client.SendEmail(ctx, email)
	if err != nil {
		panic(err)
	}

	fmt.Println("trace id:", result.TraceID())
}
