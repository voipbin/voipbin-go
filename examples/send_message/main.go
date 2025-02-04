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

	body := voipbin_client.PostMessagesJSONRequestBody(voipbin_client.PostMessagesJSONBody{
		Destinations: destinations,
		Source:       source,
		Text:         "Greetings from VoipBin!",
	})

	res, err := client.PostMessagesWithResponse(ctx, body)
	if err != nil {
		fmt.Printf("Error sending SMS message: %s\n", err)
	} else {
		fmt.Printf("Response. message_id: %s\n", *res.JSON200.Id)
	}
}
