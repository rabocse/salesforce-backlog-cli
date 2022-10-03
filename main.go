package main

import (
	"net/http"

	"github.com/rabocse/salesforce-backlog-cli/sftool"
)

func main() {

	// Getting the credentials for authentication via environment variables.
	data := sftool.EnvHandler()

	// Building Salesforce URL for authentication purposes.
	authURL := sftool.BuildURL(data.SalesforceInstance, 1)

	// Parsing the credentials.
	authPayload := sftool.CraftPayload(data.Username, data.Password, data.ClientID, data.ClientSecret, data.SecurityKey, "auth")

	// Crafting a valid HTTPS request with TLS ignore for authentication.
	authReq := sftool.CraftRequest(http.MethodPost, authURL, "no-token", authPayload)

	// Sending the request and getting a valid server response for authentication.
	authResponse := sftool.SendRequest(authReq)

	// Extracting the access token value from the server response.
	accessToken := sftool.ExtractAuthToken(authResponse)

	// CLI ...
	// cli.Cli()

	// Building the URL to query the data.
	casesURL := sftool.BuildURL(data.SalesforceInstance, 2)

	// Crafting a valid HTTPS request with TLS ignore.
	casesReq := sftool.CraftRequest(http.MethodGet, casesURL, accessToken, nil)

	// Sending the request and getting a valid server response.
	casesResponse := sftool.SendRequest(casesReq)

	// Parsing the JSON response.
	output := sftool.UnmarshalSF(casesResponse)

	// Printing the relevant info from the response.
	sftool.PrettyPrintBacklog(output)

}
