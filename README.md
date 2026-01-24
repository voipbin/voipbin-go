# voipbin-go

[![Go Reference](https://pkg.go.dev/badge/github.com/voipbin/voipbin-go.svg)](https://pkg.go.dev/github.com/voipbin/voipbin-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/voipbin/voipbin-go)](https://goreportcard.com/report/github.com/voipbin/voipbin-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

The official Go SDK for [VoIPBIN](https://voipbin.net) — build voice, SMS, and AI-powered communication apps in minutes.

## What You Can Build

- 📞 **Voice Calls** — Make, receive, and manage calls programmatically
- 💬 **SMS & Messaging** — Send and receive text messages globally
- 🤖 **AI Agents** — Build intelligent voice assistants and chatbots
- 📊 **Campaigns** — Run automated outbound calling campaigns
- 🎙️ **Recordings & Transcripts** — Record calls and get AI transcriptions
- 💻 **Conferences** — Create multi-party conference calls
- 📝 **Flows** — Design call flows with our visual builder
- 🔐 **Authentication** — Secure JWT-based login for agents

## Installation

```bash
go get github.com/voipbin/voipbin-go
```

Requires Go 1.23 or later.

## Quick Start

### Authentication

Start by logging in with your credentials to get a JWT token:

```go
package main

import (
    "context"
    "fmt"

    "github.com/voipbin/voipbin-go/gens/voipbin_client"
)

func main() {
    client, _ := voipbin_client.NewClientWithResponses("https://api.voipbin.net")

    res, err := client.PostAuthLoginWithResponse(context.Background(),
        voipbin_client.PostAuthLoginJSONRequestBody{
            Username: "your-username",
            Password: "your-password",
        })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Token: %s\n", res.JSON200.Token)
}
```

### Using Your API Key

For most API calls, use your access key:

```go
import "github.com/voipbin/voipbin-go"

client, err := voipbin.NewClient("your-access-key")
if err != nil {
    panic(err)
}
// Now use client for API calls
```

## Examples

### Send an SMS

```go
destinations := []voipbin_client.CommonAddress{
    {Target: voipbin.StringPtr("+1234567890")},
}

res, err := client.PostMessagesWithResponse(ctx, voipbin_client.PostMessagesJSONRequestBody{
    Destinations: destinations,
    Source:       voipbin_client.CommonAddress{Target: voipbin.StringPtr("+1987654321")},
    Text:         "Hello from VoIPBIN!",
})

fmt.Printf("Message ID: %s\n", *res.JSON200.Id)
```

### Make a Call

```go
destinations := []voipbin_client.CommonAddress{
    {Target: voipbin.StringPtr("+1234567890")},
}

res, err := client.PostCallsWithResponse(ctx, voipbin_client.PostCallsJSONRequestBody{
    Destinations: &destinations,
    Source:       &voipbin_client.CommonAddress{Target: voipbin.StringPtr("+1987654321")},
    FlowId:       voipbin.StringPtr("your-flow-id"),
})

for _, call := range *res.JSON200.Calls {
    fmt.Printf("Call ID: %s\n", *call.Id)
}
```

### List Your Calls

```go
res, err := client.GetCallsWithResponse(ctx, &voipbin_client.GetCallsParams{})

for _, call := range *res.JSON200.Result {
    fmt.Printf("Call: %s - Status: %s\n", *call.Id, *call.Status)
}
```

### More Examples

Check out the [examples](./examples) directory for complete, runnable programs:

| Example | Description |
|---------|-------------|
| [simple_login](./examples/simple_login) | Authenticate with username/password |
| [send_message](./examples/send_message) | Send an SMS message |
| [make_call](./examples/make_call) | Initiate an outbound call |
| [simple_client](./examples/simple_client) | List calls with pagination |

Run any example:

```bash
go run ./examples/simple_login -u your@email.com -p your-password
```

## Documentation

- 📖 [API Reference](https://api.voipbin.net/docs/) — Full endpoint documentation
- 🚀 [Getting Started Guide](https://api.voipbin.net/docs/intro.html) — Step-by-step tutorials

## Contributing

Found a bug or have a feature request? [Open an issue](https://github.com/voipbin/voipbin-go/issues)!

Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

## License

[MIT](LICENSE)
