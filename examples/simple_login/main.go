package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/spf13/cobra"
	"github.com/voipbin/voipbin-go/gens/voipbin_client"
)

// Auth endpoint is at /auth/login (not under /v1.0 prefix)
const authServerAddress = "https://api.voipbin.net"

var (
	username string
	password string
	verbose  bool
)

// debugTransport logs requests and responses when verbose mode is enabled
type debugTransport struct {
	verbose bool
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.verbose {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Printf("Failed to dump request: %s\n", err)
		} else {
			fmt.Printf("=== REQUEST ===\n%s\n", string(dump))
		}
	}

	resp, err := http.DefaultTransport.RoundTrip(req)

	if t.verbose && resp != nil {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Printf("Failed to dump response: %s\n", err)
		} else {
			fmt.Printf("=== RESPONSE ===\n%s\n", string(dump))
		}
	}

	return resp, err
}

var rootCmd = &cobra.Command{
	Use:   "simple_login",
	Short: "Login to VoIPBIN with username and password",
	Long:  "A simple example that authenticates with VoIPBIN using username/password and receives a JWT token.",
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" || password == "" {
			fmt.Println("Error: username and password are required")
			cmd.Help()
			os.Exit(1)
		}

		fmt.Printf("Attempting login for user: %s\n", username)
		fmt.Printf("Target server: %s\n", authServerAddress)

		client, err := voipbin_client.NewClientWithResponses(
			authServerAddress,
			voipbin_client.WithHTTPClient(&http.Client{
				Transport: &debugTransport{
					verbose: verbose,
				},
			}),
		)
		if err != nil {
			fmt.Printf("Failed to create client: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Client created successfully")

		ctx := context.Background()

		fmt.Println("Sending login request to /auth/login ...")
		res, err := client.PostAuthLoginWithResponse(ctx, voipbin_client.PostAuthLoginJSONRequestBody{
			Username: username,
			Password: password,
		})
		if err != nil {
			fmt.Printf("Request failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Response status: %s\n", res.Status())

		if res.JSON200 == nil {
			fmt.Printf("Login failed: %s\n", res.Status())
			if len(res.Body) > 0 {
				fmt.Printf("Response body: %s\n", string(res.Body))
			}
			os.Exit(1)
		}

		fmt.Println("")
		fmt.Println("=== Login Successful ===")
		fmt.Println("")
		fmt.Printf("Username: %s\n", res.JSON200.Username)
		fmt.Printf("Token:    %s\n", res.JSON200.Token)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "VoIPBIN username (required)")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "VoIPBIN password (required)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output (show HTTP request/response)")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
