package voipbin_test

import (
	"context"
	"fmt"
	"log"

	"github.com/voipbin/voipbin-go"
	"github.com/voipbin/voipbin-go/gens/voipbin_client"
)

// ExampleNewClient demonstrates creating a client with access key authentication.
func ExampleNewClient() {
	client, err := voipbin.NewClient("your-access-key")
	if err != nil {
		log.Fatal(err)
	}

	// Use client for API calls
	_ = client
}

// ExampleNewClientWithBasicAuth demonstrates creating a client for login endpoints.
func ExampleNewClientWithBasicAuth() {
	client, err := voipbin.NewClientWithBasicAuth("username", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Use client for authentication endpoints
	_ = client
}

// ExampleStringPtr demonstrates converting a string to a pointer for optional fields.
func ExampleStringPtr() {
	phoneNumber := voipbin.StringPtr("+1234567890")
	fmt.Printf("Type: %T\n", phoneNumber)
	// Output: Type: *string
}

// ExampleIntPtr demonstrates converting an int to a pointer for optional fields.
func ExampleIntPtr() {
	limit := voipbin.IntPtr(100)
	fmt.Printf("Type: %T\n", limit)
	// Output: Type: *int
}

// This example demonstrates listing calls with pagination.
func Example_listCalls() {
	client, err := voipbin.NewClient("your-access-key")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	params := &voipbin_client.GetCallsParams{}

	res, err := client.GetCallsWithResponse(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	if res.JSON200 != nil && res.JSON200.Result != nil {
		for _, call := range *res.JSON200.Result {
			fmt.Printf("Call ID: %s\n", *call.Id)
		}
	}
}

// This example demonstrates sending an SMS message.
func Example_sendMessage() {
	client, err := voipbin.NewClient("your-access-key")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	destinations := []voipbin_client.CommonAddress{
		{Target: voipbin.StringPtr("+1234567890")},
	}

	body := voipbin_client.PostMessagesJSONRequestBody{
		Destinations: destinations,
		Source:       voipbin_client.CommonAddress{Target: voipbin.StringPtr("+1987654321")},
		Text:         "Hello from VoIPBIN!",
	}

	res, err := client.PostMessagesWithResponse(ctx, body)
	if err != nil {
		log.Fatal(err)
	}

	if res.JSON200 != nil {
		fmt.Printf("Message sent: %s\n", *res.JSON200.Id)
	}
}

// This example demonstrates making an outbound call.
func Example_makeCall() {
	client, err := voipbin.NewClient("your-access-key")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	destinations := []voipbin_client.CommonAddress{
		{Target: voipbin.StringPtr("+1234567890")},
	}
	source := voipbin_client.CommonAddress{Target: voipbin.StringPtr("+1987654321")}
	flowID := "your-flow-id"

	body := voipbin_client.PostCallsJSONRequestBody{
		Destinations: &destinations,
		Source:       &source,
		FlowId:       &flowID,
	}

	res, err := client.PostCallsWithResponse(ctx, body)
	if err != nil {
		log.Fatal(err)
	}

	if res.JSON200 != nil && res.JSON200.Calls != nil {
		for _, call := range *res.JSON200.Calls {
			fmt.Printf("Call created: %s\n", *call.Id)
		}
	}
}
