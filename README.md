# Salesforce Backlog CLI 

Golang script for interacting with Salesforce API. 

The script is meant to be used mainly for Salesforce users instead of admins.

The drive behind this was/is to automate repetitive and tedious tasks within the technical support engineer role like:

- Reviewing backlog of cases via Salesforce GUI.
- Retrieving case information via Salesforce GUI.
- Retrieving attachments via Salesforce GUI.

To accomplish such, the script is executed as a CLI tool.

# Requirements

- Internet connection.
- API Salesforce access.
- Docker (Recommended script execution)

NOTE: Source code and compiled binary are available in case of not having/preferring Docker. However, the execution is recommended via Docker.


# TLDR (Execution with Docker)

1. Get the CLI tool.
```
docker run -it rabocse/salesforce-backlog-cli
```
<br/>

2. In the container CLI, proceed to paste Salesforce Credentials: email, password, salesforce instance, client ID, client Secret and security key. 
   
```
export EMAIL=rabocse@mydomain.com
export PASS=MyFakePassword123
export SF=myfake.sf.instance.salesforce.com
export CLID=xxxxxxxyyyyyyyyyyaaaaaaabbbbbbbbdddddddddddd22211111
export CLSE=11111112222222333333344444aaaaaccccc1112222222
export SECK=BAD23XXXXXXXXFFF
```
<br/>

3. Get your backlog of cases:
```
./main
```

# Contents

- [Salesforce Backlog CLI](#salesforce-backlog-cli)
- [Requirements](#requirements)
- [TLDR (Execution with Docker)](#tldr-execution-with-docker)
- [Contents](#contents)
- [Current State](#current-state)
- [Script's Data Flow](#scripts-data-flow)
- [Caveats](#caveats)
- [Progress and Roadmap](#progress-and-roadmap)



# Current State

The tool successfully requests an access token to the Salesforce Instance to authenticate and then downloads the data from sObjects/case resource ("/services/data/v55.0/sobjects/case").

First, the user must set the needed credentials to authenticate against the Salesforce API. For such, environment variables are passed in the container CLI.

Once, the environment variables are set, the tool is ready to be executed:

```
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

```

# Script's Data Flow

- The script reads the environment variables (EMAIL, PASS, SF, CLID, CLSE, SECK) from the user's terminal.

- Such environment variables are used to authenticate against the Salesforce instance and get an access token.

- Once the access token is properly downloaded and parsed, then it is used in a HTTP header for the data request.

- The listview (Salesforce resource) is queried and successfully unmarshalled to then be printed in a table format.


# Caveats

- The [listview ID](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_listviews.htm) is currently hardcoded within sftool.BuildURL function. 
  
- The Salesforce API version is also hardcoded to v55.0. 



# Progress and Roadmap

- Input for the script to be accepted via enviroment variables. __[DONE]__
- Parse JSON output and extract only the "access token" value. __[DONE]__
- Pass the access token value to next section of the script. __[DONE]__
- Avoid the usage of external tools (jq and/or grep), build the presentation of data in the source code. (Table format) __[DONE]__
- Create "Go modules" and share the interim state of the script. __[DONE]__
- Containerize the application. __[DONE]__
- Modify the downloaded resource (currently sObject/case) to a resource that provides the list of active cases from the engineer.
- Allow the user to specify a case ID to get additional information about it.
- Get the attachment from Salesforce:
    - Direct attachments from Salesforce.
    - Attachments from third-party integrated tool like S-Drive.
- Add logic business or a feature to specify the listview ID to be retrieved. (See caveat for reference)

<br/>

__NOTE:__ Roadmap is subject to changes.



