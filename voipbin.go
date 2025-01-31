package voipbin

import (
	"net/http"

	"github.com/voipbin/voipbin-go/gens/voipbin_client"
)

const (
	defaultServerAddress = "https://api.voipbin.net/v1.0"
)

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

func StrPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}
