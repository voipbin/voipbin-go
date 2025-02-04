# voipbin-go
A Go client for the [VoIPBIN API](https://api.voipbin.net/docs/intro.html), enabling developers to easily interact with VoIPBIN's cloud-based communication services.

## 🚀 Installation

Install `voipbin-go` using:

```sh
go get github.com/voipbin/voipbin-go
```

## 🌍 Quickstart
Try sending yourself an SMS message by pasting the following code example into a send_message in the examples directory where you installed voipbin-go. Be sure to update the accesskey and phone number value from your voipbin account.

```go
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
			Target: voipbin.StrPtr("<your phone number here>"),
		},
	}
	source := voipbin_client.CommonAddress{
		Target: voipbin.StrPtr("+1234567892"),
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
```

## 📞 Making a Call
You can initiate a call using the voipbin_client.PostCallsJSONRequestBody() method:
```go
	destinations := []voipbin_client.CommonAddress{
		{
			Target: voipbin.StrPtr("+1234567890"),
		},
	}
	source := voipbin_client.CommonAddress{
		Target: voipbin.StrPtr("+1234567892"),
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
```

## Getting help
If you need help installing or using the library, please check the voipbin's api documentation first, and file a support ticket if you don't find an answer to your question.

If you've instead found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!

* https://api.voipbin.net/docs/