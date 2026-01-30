// Package voipbin provides a Go client for the VoIPBIN API.
//
// VoIPBIN is a cloud-based communication platform that enables you to build
// voice, SMS, and AI-powered applications. This SDK provides type-safe bindings
// generated from the OpenAPI specification.
//
// # Capabilities
//
//   - Voice Calls: Make, receive, and manage calls programmatically
//   - SMS & Messaging: Send and receive text messages globally
//   - AI Agents: Build intelligent voice assistants and chatbots
//   - Campaigns: Run automated outbound calling campaigns
//   - Recordings & Transcripts: Record calls and get AI transcriptions
//   - Conferences: Create multi-party conference calls
//
// # Authentication
//
// The SDK supports two authentication methods:
//
//   - Access Key: Use [NewClient] for most API operations
//   - Basic Auth: Use [NewClientWithBasicAuth] for agent login endpoints
//
// # Usage
//
// Create a client, build a request using generated types, and call the API:
//
//	client, err := voipbin.NewClient("your-access-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	res, err := client.GetCallsWithResponse(ctx, &voipbin_client.GetCallsParams{})
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	for _, call := range *res.JSON200.Result {
//	    fmt.Printf("Call: %s\n", *call.Id)
//	}
//
// For complete API documentation, see https://api.voipbin.net/docs/
package voipbin

import (
	"net/http"

	"github.com/voipbin/voipbin-go/gens/voipbin_client"
)

const (
	defaultServerAddress = "https://api.voipbin.net/v1.0"
)

// NewClient creates a new VoIPBIN API client using access key authentication.
//
// This is the primary way to create a client for most API operations including
// making calls, sending messages, managing campaigns, and more.
//
// The access key is automatically appended to all requests as a query parameter.
//
// See examples/simple_client, examples/send_message, and examples/make_call
// for complete runnable examples.
func NewClient(accesskey string) (voipbin_client.ClientWithResponsesInterface, error) {
	res, err := voipbin_client.NewClientWithResponses(defaultServerAddress, withAccessKey(accesskey))
	if err != nil {
		return nil, err
	}

	return res, nil
}

type customTransport struct {
	accessKey string
}

// RoundTrip is the method that modifies the URL to add the accesskey.
func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the URL to modify it.
	newURL := *req.URL
	query := newURL.Query()
	query.Set("accesskey", t.accessKey)
	newURL.RawQuery = query.Encode()

	// Clone the original request with the new URL.
	newReq := req.Clone(req.Context())
	newReq.URL = &newURL

	// Use the default HTTP client to send the modified request.
	return http.DefaultClient.Do(newReq)
}

func withAccessKey(accessKey string) voipbin_client.ClientOption {
	return func(c *voipbin_client.Client) error {
		// Wrap the existing transport with the custom transport.
		c.Client = &http.Client{
			Transport: &customTransport{
				accessKey: accessKey,
			},
		}
		return nil
	}
}

// NewClientWithBasicAuth creates a new VoIPBIN API client using Basic Authentication.
//
// This is used for agent login endpoints that require username/password authentication,
// such as the /auth/login endpoint which returns a JWT token.
//
// See examples/simple_login for a complete runnable example.
func NewClientWithBasicAuth(username, password string) (voipbin_client.ClientWithResponsesInterface, error) {
	res, err := voipbin_client.NewClientWithResponses(defaultServerAddress, withBasicAuth(username, password))
	if err != nil {
		return nil, err
	}

	return res, nil
}

type basicAuthTransport struct {
	username string
	password string
}

// RoundTrip adds Basic Auth header to the request.
func (t *basicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := req.Clone(req.Context())
	newReq.SetBasicAuth(t.username, t.password)
	return http.DefaultClient.Do(newReq)
}

func withBasicAuth(username, password string) voipbin_client.ClientOption {
	return func(c *voipbin_client.Client) error {
		c.Client = &http.Client{
			Transport: &basicAuthTransport{
				username: username,
				password: password,
			},
		}
		return nil
	}
}

// StringPtr returns a pointer to the given string value.
//
// This helper is useful when working with the generated API types, which use
// pointers for optional fields. For example:
//
//	address := voipbin_client.CommonAddress{
//	    Target: voipbin.StringPtr("+1234567890"),
//	}
//
// See examples/send_message and examples/make_call for usage in context.
func StringPtr(s string) *string {
	return &s
}

// IntPtr returns a pointer to the given int value.
//
// This helper is useful when working with the generated API types, which use
// pointers for optional fields.
func IntPtr(i int) *int {
	return &i
}
