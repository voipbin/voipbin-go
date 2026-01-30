# SDK Documentation Design

## Problem

The pkg.go.dev page for `github.com/voipbin/voipbin-go` shows "Documentation not displayed due to license restrictions" with "License: None detected". The README mentions MIT but no LICENSE file exists.

## Solution

Add LICENSE file + comprehensive Go doc comments + testable examples.

## Changes

### 1. LICENSE File (NEW)

Standard MIT license with copyright "2026 VoIPBIN". Required for pkg.go.dev to display documentation.

### 2. Package Documentation (voipbin.go)

Add package-level doc comment covering:
- What the SDK does (Go client for VoIPBIN API)
- Key capabilities (voice calls, SMS, AI agents, campaigns, recordings)
- Authentication methods (access key vs Basic Auth)
- Quick usage pattern (create client → build request → call API)
- Link to full API docs

### 3. Function Documentation (voipbin.go)

Add doc comments to exported functions:

**NewClient(accesskey string)**
- Creates API client with access key authentication
- References: examples/simple_client, examples/send_message, examples/make_call

**NewClientWithBasicAuth(username, password string)**
- Creates API client with Basic Auth for login endpoints
- References: examples/simple_login

**StringPtr(s string) / IntPtr(i int)**
- Helpers to convert values to pointers for optional fields
- Usage shown in: examples/send_message, examples/make_call

### 4. Testable Examples (example_test.go - NEW)

Create Example* functions that appear on pkg.go.dev and run during `go test`:

| Function | Demonstrates |
|----------|-------------|
| ExampleNewClient | Basic client creation with access key |
| ExampleNewClientWithBasicAuth | Client creation for login endpoints |
| ExampleStringPtr | Converting string to pointer |
| Example_sendMessage | Full SMS sending flow |
| Example_listCalls | Listing calls with pagination |

## File Structure

```
voipbin-go/
├── LICENSE                 # NEW: MIT license (2026)
├── voipbin.go              # MODIFY: Add package doc + function doc comments
├── example_test.go         # NEW: Testable Example* functions
```

## Expected Result

After pushing, pkg.go.dev will:
- Detect the MIT license
- Display the package overview from doc comments
- Show the Examples section with runnable code
- Link to the generated voipbin_client sub-package
