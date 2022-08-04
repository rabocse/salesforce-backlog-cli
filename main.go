package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

/* The current execution is succesful:

❯ ./main -user=$NAME -pass=$PASS -sf=$SF -clid=$CLID -clse=$CLSE -seck=$SECK
[Ommited output due confidentialy of info]

*/

const method string = "POST"

// flagsHander parses the flags passed by the user via CLI
func flagsHandler() (s, u, p, ci, cs, sk string) {

	// Requesting flags to user via CLI.
	// NOTE: flag.String returns a pointer.
	sfi := flag.String("sf", " ", "Salesforce Instance, e.g  https://mycompany.salesforce.com")
	user := flag.String("user", " ", "Salesforce Username, e.g rabocse@mycompany.com")
	pass := flag.String("pass", " ", "Salesforce Password, e.g mysecurefakepassword123")
	clid := flag.String("clid", " ", "Salesforce Client ID (This should be provided by your Salesforce Admin)")
	clse := flag.String("clse", " ", "Salesforce Client Secret (This should be provided by your Salesforce Admin)")
	seck := flag.String("seck", " ", "Salesforce Security Key (This should be provided by your Salesforce Admin)")

	// Execute the command-line parsing
	flag.Parse()

	// Convert the string pointer to a string
	s = *sfi
	u = *user
	p = *pass
	ci = *clid
	cs = *clse
	sk = *seck

	return s, u, p, ci, cs, sk

}

//  buildURL returns a valid string URL
func buildURL(salesforceInstance string) string {

	// Define the components for the HTTP Request.
	const protocol string = "https://"
	const resource string = "/services/oauth2/token"

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)

	return url
}

// craftPayload prepares the credentials to be added as payload to a valid HTTP(s) request.
func craftPayload(userValue, passwordValue, clientIDvalue, clientSecretvalue, securityKeyvalue string) io.Reader {

	c := struct {
		Username     string
		Password     string
		GrantType    string
		ClientID     string
		ClientSecret string
		SecurityKey  string
	}{
		Username:     userValue,
		Password:     passwordValue,
		GrantType:    "password",
		ClientID:     clientIDvalue,
		ClientSecret: clientSecretvalue,
		SecurityKey:  securityKeyvalue,
	}

	// Concatenate to build the payload
	concatenatedPayload := fmt.Sprintf("grant_type=%s&client_id=%s&client_secret=%s&username=%s&password=%s%s", c.GrantType, c.ClientID, c.ClientSecret, c.Username, c.Password, c.SecurityKey) // concatenatedPayload is a string (non encoded)

	// Convert to *strings.Reader
	p := strings.NewReader(concatenatedPayload)

	return p
}

// craftRequest prepares a valid HTTP request with a POST method and the specified URL and payload.
func craftRequest(m string, u string, p io.Reader) *http.Request {

	// Build the request (req) with the previous components
	req, err := http.NewRequest(m, u, p)

	if err != nil {
		fmt.Println(err)
	}

	// Header to specify that our request sends plain text format.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req

}

// sendRequest executes the so far crafted Request.
func sendRequest(r *http.Request) string {

	// Make the Go client to ignore the TLS verification
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transCfg}

	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	b := string(body)

	return b

}

func main() {

	// Values are passed via CLI
	salesforceInstance, username, password, clientID, clientSecret, SecurityKey := flagsHandler()

	// Cluster URL is built.
	url := buildURL(salesforceInstance)

	// Credentials are parsed to be payload.
	payload := craftPayload(username, password, clientID, clientSecret, SecurityKey) // <==== Currently working on this function.

	// Crafting a valid HTTPS request with TLS ignore.
	req := craftRequest(method, url, payload)

	// Sending the request and getting a valid authToken
	authToken := sendRequest(req)

	// Printing the authentication token
	fmt.Println(authToken)
}
