package main

import (
	"context"
	"fmt"

	"github.com/voipbin/voipbin-go"
	"github.com/voipbin/voipbin-go/gens/voipbin_client"
)

func main() {
	client, err := voipbin.NewClient("<your api accesskey here>")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	params := &voipbin_client.GetCallsParams{}
	tmp, err := client.GetCallsWithResponse(ctx, params)
	if err != nil {
		panic(err)
	}

	if tmp.JSON200 == nil {
		panic("unexpected nil response")
	}

	fmt.Printf("Next page token: %s\n", *tmp.JSON200.NextPageToken)
	for i, c := range *tmp.JSON200.Result {
		fmt.Printf("Call %d: %v\n", i, *c.Id)
	}

}
