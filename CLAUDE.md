# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

voipbin-go is a Go client library for the VoIPBIN API. It provides type-safe bindings generated from an OpenAPI specification for interacting with VoIPBIN's cloud-based communication services (calls, SMS, AI agents, campaigns, etc.).

**API Docs**: https://api.voipbin.net/docs/intro.html

## Commands

### Code Generation

The client code is auto-generated from an OpenAPI spec. To regenerate:

```bash
go generate ./openapi/config_client/
```

**Note**: The OpenAPI spec source is at `../monorepo/bin-openapi-manager/openapi/openapi.yaml` (relative to this repo).

### Build and Test

```bash
go build ./...                    # Build all packages
go test ./...                     # Run all tests
go mod tidy                       # Clean up dependencies
go mod vendor                     # Update vendor directory
```

### Running Examples

```bash
go run ./examples/send_message/   # SMS sending example
go run ./examples/make_call/      # Call initiation example
go run ./examples/simple_client/  # API listing example
```

## Architecture

### Key Files

- **`voipbin.go`** - Main package wrapper providing `NewClient()` and helper functions
- **`gens/voipbin_client/gen.go`** - Auto-generated OpenAPI client (~40K lines, DO NOT EDIT directly)
- **`openapi/config_client/`** - Code generation configuration and build directive

### Authentication Flow

The client uses query parameter authentication via a custom HTTP transport:

```
NewClient(accesskey)
  → creates ClientWithResponses
  → wraps HTTP transport with customTransport
  → customTransport.RoundTrip() injects ?accesskey=... on every request
```

Default API endpoint: `https://api.voipbin.net/v1.0`

### Usage Pattern

```go
// 1. Create client
client, err := voipbin.NewClient("<api-key>")

// 2. Build request using generated types (note: uses pointers)
body := voipbin_client.PostMessagesJSONRequestBody{
    Destinations: []voipbin_client.CommonAddress{{Target: voipbin.StringPtr("+1...")}},
    Source:       voipbin_client.CommonAddress{Target: voipbin.StringPtr("+1...")},
    Text:         "Hello",
}

// 3. Call API (all methods end with WithResponse)
res, err := client.PostMessagesWithResponse(ctx, body)

// 4. Access typed response
if res.JSON200 != nil {
    messageID := *res.JSON200.Id
}
```

### Helper Functions

- `voipbin.StringPtr(s string) *string` - Convert string to pointer (needed for optional fields)
- `voipbin.IntPtr(i int) *int` - Convert int to pointer

## Code Generation Details

Uses **oapi-codegen** to generate Go client from OpenAPI 3.0 spec:

- Generator: `github.com/oapi-codegen/oapi-codegen/v2`
- Config: `openapi/config_client/config.generate.yaml`
- Output: `gens/voipbin_client/gen.go`

The generated code includes:
- All request/response struct types with JSON tags
- Enum constants for API values
- `ClientWithResponsesInterface` with typed `*WithResponse()` methods
- HTTP client with `ClientOption` pattern for customization

## Dependency Management

Uses Go modules with vendoring. After modifying dependencies:

```bash
go mod tidy
go mod vendor
```

## Commit Message Format

**Title (first line):**
VOIP-[ticket-number]-brief-description-of-change

or (when no JIRA ticket)
NOJIRA-brief-description-of-change

**Body (subsequent lines):**
List each affected project with specific changes:

- project-name: Specific change or feature
- project-name: Another specific change

**Example:**
```
NOJIRA-Add-auth-login-openapi-spec

Add /auth/login endpoint to OpenAPI specification with AuthLoginResponse schema.

- bin-openapi-manager: Add Auth tag to openapi.yaml
- bin-openapi-manager: Add AuthLoginResponse schema in components/schemas
- bin-openapi-manager: Create paths/auth/login.yaml with POST endpoint spec
- bin-api-manager: Implement PostAuthLogin handler using generated types
- bin-api-manager: Add auth_test.go with login test cases
```

**Rules:**
1. Title MUST match the branch name exactly
2. Always list affected projects with bullet points
3. Be specific about what changed in each project
4. Keep title concise (use underscores for spaces)
5. Use present tense
6. Use dashes (-) for bullet points
7. Add narrative summary for significant changes
