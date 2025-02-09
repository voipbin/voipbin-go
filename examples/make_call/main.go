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

	destinations := []voipbin_client.CommonAddress{
		{
			Target: voipbin.StringPtr("<your phone number here>"),
		},
	}
	source := voipbin_client.CommonAddress{
		Target: voipbin.StringPtr("+1234567892"),
	}
	flowID := "<your flow id here>"

	body := voipbin_client.PostCallsJSONRequestBody(voipbin_client.PostCallsJSONBody{
		Destinations: &destinations,
		FlowId:       &flowID,
		Source:       &source,
	})

	res, err := client.PostCallsWithResponse(ctx, body)
	if err != nil {
		panic(err)
	}

	for i, c := range *res.JSON200.Calls {
		fmt.Printf("Created Call %d: %v\n", i, *c.Id)
	}
	for i, c := range *res.JSON200.Groupcalls {
		fmt.Printf("Created Groupcall %d: %v\n", i, *c.Id)
	}
}
