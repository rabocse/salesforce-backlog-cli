package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*

The current execution is succesful. First the user must set the expected enviroment variables on the local terminal. For example:

---
❯ export EMAIL=rabocse@mydomain.com
export PASS=MyFakePassword123
export SF=myfake.sf.instance.salesforce.com
export CLID=xxxxxxxyyyyyyyyyyaaaaaaabbbbbbbbdddddddddddd22211111
export CLSE=11111112222222333333344444aaaaaccccc1112222222
export SECK=BAD23XXXXXXXXFFF
---

And then proceed to execute:

❯ ./accessToken
[Ommited output due confidentialy of info]

*/

const method string = "POST"

func envHandler() (sfi, user, pass, clid, clse, seck string) {

	// Get needed enviroment variables: EMAIL, PASS, SF, CLID, CLSE, SECK.
	sfi = os.Getenv("SF")
	user = os.Getenv("EMAIL")
	pass = os.Getenv("PASS")
	clid = os.Getenv("CLID")
	clse = os.Getenv("CLSE")
	seck = os.Getenv("SECK")

	return sfi, user, pass, clid, clse, seck

}

/*

//  buildURL returns a valid string URL
func buildAuthURL(salesforceInstance string) string {

	// Define the components for the HTTP Request.
	const protocol string = "https://"
	const resource string = "/services/oauth2/token"

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)

	return url
}

//  buildURL returns a valid string URL
func buildSFURL(salesforceInstance string) string {

	// Define the components for the HTTP Request.
	const protocol string = "https://"
	const resource string = "/services/data/v55.0/sobjects/case" // TODO: This is the resource to modify depending on what needs to be downloaded. For the time being, backlog cases.

	// Concatenate to build the URL
	url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)

	return url
}

*/

// buildURL builds any URL resource (API resource). No need of duplicate functions per each resource.
func buildURL(salesforceInstance string, resource int) string {

	const protocol string = "https://"
	var apiResources int = resource

	switch apiResources {

	case 1:

		const resource string = "/services/oauth2/token"
		// Concatenate to build the URL
		url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)
		return url

	case 2:

		const resource string = "/services/data/v55.0/sobjects/case"

		// Concatenate to build the URL
		url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)

		return url

	}

	invalidResource := fmt.Sprintf("The specified resource (%d) did not match any available option.", resource)

	return invalidResource
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

	// Header to specify that our request sends urlencoded format.
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

func extractAuthToken(r string) string {

	type response struct {
		AccessToken string `json:"access_token"`
		InstanceURL string `json:"instance_url"`
		Id          string `json:"id"`
		TokenType   string `json:"token_type"`
		IssuedAt    string `json:"issued_at"`
		Signature   string `json:"signature"`
	}

	rByte := []byte(r)

	var serverResponse response
	json.Unmarshal(rByte, &serverResponse)

	return serverResponse.AccessToken

}

func main() {

	// Values are passed via CLI
	salesforceInstance, username, password, clientID, clientSecret, SecurityKey := envHandler()

	// Builds Salesforce URL
	authURL := buildURL(salesforceInstance, 1)

	// Credentials are parsed to be payload.
	payload := craftPayload(username, password, clientID, clientSecret, SecurityKey)

	// Crafting a valid HTTPS request with TLS ignore.
	req := craftRequest(method, authURL, payload)

	// Sending the request and getting a valid server response
	response := sendRequest(req)

	accessToken := extractAuthToken(response)

	// Printing the authentication token value
	fmt.Println()
	fmt.Println(accessToken)

	backlogURL := buildURL(salesforceInstance, 2) // TODO: Next... I need to take the backlogURL and craft the request with the new payload for such.

	// Printing the backlog URL
	fmt.Println()
	fmt.Println(backlogURL)

	// TODO: The amount of functions are getting doubled. Refactor will be needed to shorten the source code.
}
