package sftool

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// envHandler gets the needed enviroment variables: EMAIL, PASS, SF, CLID, CLSE, SECK.
func EnvHandler() (sfi, user, pass, clid, clse, seck string) {

	sfi = os.Getenv("SF")
	user = os.Getenv("EMAIL")
	pass = os.Getenv("PASS")
	clid = os.Getenv("CLID")
	clse = os.Getenv("CLSE")
	seck = os.Getenv("SECK")

	return sfi, user, pass, clid, clse, seck

}

// buildURL builds any URL resource (API resource). No need of duplicate functions per each resource.
func BuildURL(salesforceInstance string, resource int) string {

	const protocol string = "https://"
	var apiResources int = resource

	switch apiResources {

	case 1:

		const resource string = "/services/oauth2/token"
		// Concatenate to build the URL
		url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)
		return url

	case 2:
		// Resource: listview called "My Cases"
		const resource string = "/services/data/v55.0/sobjects/case/listviews/00BE0000004x68gMAA/results"
		// const resource string = "/services/data/v55.0/sobjects/case/listviews/xxxxxxxxxx/results" // Verifying if it works with other listviews. Replace xxxx for the ID listview.

		// Concatenate to build the URL
		url := fmt.Sprintf("%s%s%s", protocol, salesforceInstance, resource)

		return url

	}

	invalidResource := fmt.Sprintf("ERROR: The specified resource (%d) did not match any available option.", resource)

	return invalidResource
}

// craftPayload prepares the payload used in the http requests.
func CraftPayload(userValue, passwordValue, clientIDvalue, clientSecretvalue, securityKeyvalue string, purpose string) io.Reader {

	var payloadPurpose string = purpose

	switch payloadPurpose {

	case "auth":

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

	case "crud":

		// For future CRUD operations.

	}

	return nil // At this point, "auth" must be the only purpose specified, the authentication is the only carrying payload.

}

// craftRequest crafts a valid HTTP request with the passed http.Method, url(u), token(t) and payload(p).
func CraftRequest(m string, u string, t string, p io.Reader) *http.Request {

	var requestPurpose string = m

	switch requestPurpose {

	case "POST":

		if t == "no-token" { // POST for Authentication

			// Build the request (req) with the previous components
			req, err := http.NewRequest(m, u, p)

			if err != nil {
				fmt.Println(err)
			}

			// Header to specify that our request sends urlencoded format.
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			return req

		} else { // POST for Writing Operation

			fmt.Println("Writing Operations are not supported yet. Feel free to contribute at https://github.com/rabocse/salesforce-backlog-cli   :)")

		}

	case "GET":

		// Build the request (req) with the previous components
		req, err := http.NewRequest(m, u, p)

		if err != nil {
			fmt.Println(err)
		}

		// Header to specify that our request sends urlencoded format.
		req.Header.Add("Authorization", t)

		return req

	}

	return nil

}

// sendRequest executes the so far crafted Request.
func SendRequest(r *http.Request) string {

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

// extractAuthToken extracts the access_token value from the response sent by the server
func ExtractAuthToken(r string) string {

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

	token := fmt.Sprintf("Bearer %s", serverResponse.AccessToken)

	return token
}

// Data structure from listview response.
type listview struct {
	Columns []Columns `json:"columns"`
	// DeveloperName string    `json:"developerName"`
	// Done          bool      `json:"done"`
	// ID            string    `json:"id"`
	// Label         string    `json:"label"`
	Records []Records `json:"records"`
	// Size          int       `json:"size"`
}

type Columns struct {
	FieldNameOrPath string `json:"fieldNameOrPath"`
	Value           string `json:"value"`
}
type Records struct {
	Columns []Columns `json:"columns"`
}

// unmarshalSF unmarshals the listview response from Salesforce and returns a map.
func UnmarshalSF(cr string) map[int][]string {

	// Create a variable of listview type and unmarshal caseResonse on it.
	res := listview{}
	json.Unmarshal([]byte(cr), &res)

	majorMap := make(map[int][]string)
	mySlice := make([]string, 0)
	for k, _ := range res.Records {
		for x, _ := range res.Records[k].Columns {
			mySlice = append(mySlice, res.Records[k].Columns[x].Value)
		}
		// majorMap[res.Records[k].Columns[0].Value] = mySlice
		majorMap[k] = mySlice
		mySlice = nil
	}

	return majorMap

}

func PrettyPrintBacklog(output map[int][]string) {

	// Create the title header of the table
	title := tablewriter.NewWriter(os.Stdout)
	title.SetHeader([]string{"SALESFORCE BACKLOG"})
	title.SetRowLine(true) // Enable row line
	title.Render()

	// Create the columns of the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{" # ", "CASE NUMBER", "CONTACT NAME", "SUBJECT", "SEVERITY", "STATUS", "ENVIRONMENT"})

	// Iterate over the map and get the keys(they will not be ordered) stored in the  slice (keys).
	keys := make([]int, 0)
	for k, _ := range output {
		keys = append(keys, k)
	}

	// Order keys in the slice.
	sort.Ints(keys)

	// Iterate over the ordered slice (keys) to get an ordered value from the map output.
	for _, k := range keys {

		// counter is used for the " # " column from the table. It needs to be incremented, otherwise it starts with zero (0).
		// Also, since counter is an integer, it needs to be converted to string to be later appended.
		counter := k + 1
		part1 := []string{strconv.Itoa(counter)}

		// outputSlice containes the full value from the output map[int][]string.
		// But, we are only interested in the first 7 elements (0-6) to display in the table.
		outputSlice := output[k]
		part2 := outputSlice[0:6]

		// full is created to then be displayed.
		full := append(part1, part2...)
		table.Append(full)
	}

	table.SetRowLine(true) // Enable row line
	table.Render()

}
