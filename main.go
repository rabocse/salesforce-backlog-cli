package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const method string = "POST"

// flagsHander parses the flags passed by the user via CLI
func flagsHandler() (s, u, p string) {

	// Requesting flags to user via CLI.
	// NOTE: flag.String returns a pointer.
	sfi := flag.String("sf", " ", "Salesforce Instance, e.g  https://mycompany.salesforce.com")
	user := flag.String("user", " ", "Salesforce Username, e.g rabocse@mycompany.com")
	pass := flag.String("pass", " ", "Salesforce Password, e.g mysecurefakepassword123")

	// Execute the command-line parsing
	flag.Parse()

	// Convert the string pointer to a string
	s = *sfi
	u = *user
	p = *pass

	return s, u, p

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
func craftPayload(userValue, passwordValue string) io.Reader { //TODO: Modify this function to prepare the body of the request. At the end, an io.Reader needs to be returnedt so it can be processed by next function (http.NewRequest)

	type credentials struct {
		Username     string
		Password     string
		ClientID     string
		ClientSecret string
		SecurityKey  string
	}

	// Concatenate to build the payload
	// concatenatedPayload := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource) // concatenatedPayload is a string

	// payload := strings.NewReader("grant_type=xxxxxxxxx&client_id=xxxxxxxxxxxxxxxxx&client_secret=xxxxxxxxxxxxxxxxxx=rabocse%40mycompany.com&password=myfakepassword123myfakesecurityId") // Do I need string.NewReader?

	// convertedconcatenatedPayload := []byte(concatenatedPayload) // now concatenatedPayload is a slice of Bytes called convertedconcatenatedPayload

	// p := bytes.NewReader(convertedconcatenatedPayload) // so convertedconcatenatedPayload needs to be converted io.Reader to be accepted by next function (http.NewRequest)

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
	req.Header.Add("Content-Type", "text/plain")

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
	salesforceInstance, username, password := flagsHandler()

	// Cluster URL is built.
	url := buildURL(salesforceInstance)

	// Credentials are parsed to be payload.
	payload := craftPayload(username, password) // <==== Currently working on this function.

	// Crafting a valid HTTPS request with TLS ignore.
	req := craftRequest(method, url, payload)

	// Sending the request and getting a valid authToken
	authToken := sendRequest(req)

	// Printing the authentication token
	fmt.Println(authToken)
}
