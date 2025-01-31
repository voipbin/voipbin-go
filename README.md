# voipbin-go
A Go client for the [VoIPBIN API](https://api.voipbin.net/docs/intro.html), enabling developers to easily interact with VoIPBIN's cloud-based communication services.

## 🚀 Installation

Install `voipbin-go` using:

```sh
go get github.com/voipbin/voipbin-go
```

## 🌍 Quickstart
Try sending yourself an SMS message by pasting the following code example into a sendsms.go file in the same directory where you installed twilio-go. Be sure to update the accountSid, authToken, and from phone number with values from your Twilio account. The to phone number can be your own mobile phone number.

```go
package main

import (
	"context"
	"fmt"
	"voipbin-go"
	"voipbin-go/gens/voipbin_client"
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
```

## 📞 Making a Call
You can initiate a call using the voipbin_client.PostCallsJSONRequestBody() method:
```go
	destinations := []voipbin_client.CommonAddress{
		{
			Target: voipbin.StrPtr("+1234567890"),
		},
		{
			Target: voipbin.StrPtr("+1234567891"),
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
If you need help installing or using the library, please check the Twilio Support Help Center first, and file a support ticket if you don't find an answer to your question.

If you've instead found a bug in the library or would like new features added, go ahead and open issues or pull requests against this repo!
