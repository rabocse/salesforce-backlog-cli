package main

import (
	"net/http"

	"github.com/rabocse/salesforce-backlog-cli/sftool"
)

/*

The current execution is successful. First the user must set the expected environment variables on the local terminal. For example:

---
❯ export EMAIL=rabocse@mydomain.com
export PASS=MyFakePassword123
export SF=myfake.sf.instance.salesforce.com
export CLID=xxxxxxxyyyyyyyyyyaaaaaaabbbbbbbbdddddddddddd22211111
export CLSE=11111112222222333333344444aaaaaccccc1112222222
export SECK=BAD23XXXXXXXXFFF
---

And then proceed to execute:

❯ ./main
+--------------------+
| SALESFORCE BACKLOG |
+--------------------+
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|  #  | CASE NUMBER |  CONTACT NAME  |            SUBJECT             |   SEVERITY    | STATUS |         ENVIRONMENT         |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   1 |     1234567 | Dexter         | [AMER] Dee Dee Pushed some     | Sev1 (High)   | Open   | Dexter's Lab                |
|     |             |                | Config and Services Now Are    |               |        |                             |
|     |             |                | Down                           |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   2 |     1234568 | Johnny Bravo   | [AMER] Need Her Contact Number | Sev1 (High)   | Open   | Johnny's Mom's House        |
|     |             |                | ASAP!!!                        |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   3 |     7654321 | Dexter         | [AMER] Quantum Scaling         | Sev3 (Normal) | Open   | Dexter's Lab                |
|     |             |                | Sometimes Fails                |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   4 |     1122334 | Hyoga C.       | [EMEA] Frozen Container is     | Sev3 (Normal) | Open   | Knights of the Zodiac       |
|     |             |                | Leaking Memory                 |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   5 |     9876543 | Master Roshi   | [APAC] KameHouse Service Not   | Sev3 (Normal) | Open   |  DBZ                        |
|     |             |                | Accesible After Changing Cert  |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   6 |     1234566 | Kakashi Hatake | [APAC] Unable to Deploy        | Sev4 (Normal) | Open   | Konoha Leaf Village         |
|     |             |                | Rinnegan Service. Stuck In MS  |               |        |                             |
|     |             |                | Phase 2                        |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   7 |     1234555 | Aizen Sosuke.  | [APAC] PROACTIVE: Upgrade      | Sev3 (Low)    | Open   | Gotei 13 Inc                |
|     |             |                | Shinigami's Cluster to same    |               |        |                             |
|     |             |                | version than Espada's Cluster  |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+
|   8 |     7777777 | Tony Stark     | [EMEA] Ultron's service does   | Sev3 (Low)    | Open   | Stark Labs (Sokovia Center) |
|     |             |                | not work as expected.          |               |        |                             |
+-----+-------------+----------------+--------------------------------+---------------+--------+-----------------------------+

*/

func main() {

	// Getting the credentials for authentication via environment variables.
	salesforceInstance, username, password, clientID, clientSecret, SecurityKey := sftool.EnvHandler()

	// Building Salesforce URL for authentication purposes.
	authURL := sftool.BuildURL(salesforceInstance, 1)

	// Parsing the credentials.
	authPayload := sftool.CraftPayload(username, password, clientID, clientSecret, SecurityKey, "auth")

	// Crafting a valid HTTPS request with TLS ignore for authentication.
	authReq := sftool.CraftRequest(http.MethodPost, authURL, "no-token", authPayload)

	// Sending the request and getting a valid server response for authentication.
	authResponse := sftool.SendRequest(authReq)

	// Extracting the access token value from the server response.
	accessToken := sftool.ExtractAuthToken(authResponse)

	// Building the URL to query the data.
	casesURL := sftool.BuildURL(salesforceInstance, 2)

	// Crafting a valid HTTPS request with TLS ignore.
	casesReq := sftool.CraftRequest(http.MethodGet, casesURL, accessToken, nil)

	// Sending the request and getting a valid server response.
	casesResponse := sftool.SendRequest(casesReq)

	// Parsing the JSON response.
	output := sftool.UnmarshalSF(casesResponse)

	// Printing the relevant info from the response.
	sftool.PrettyPrintBacklog(output)

}
