# salesforce-backlog-cli (dev3.0 branch)



## Current State

The script successfully requests an access token to the Salesforce Instace to authenticate and then downloads the data from the sObjects/case resource ("/services/data/v55.0/sobjects/case").

First the user must set the expected enviroment variables on the local terminal. For example:

```
❯ export EMAIL=rabocse@mydomain.com
export PASS=MyFakePassword123
export SF=myfake.sf.instance.salesforce.com
export CLID=xxxxxxxyyyyyyyyyyaaaaaaabbbbbbbbdddddddddddd22211111
export CLSE=11111112222222333333344444aaaaaccccc1112222222
export SECK=BAD23XXXXXXXXFFF
```

And then proceed to execute:

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

- The script reads the enviroment variables (EMAIL, PASS, SF, CLID, CLSE, SECK) from the user's terminal and uses them as input.

- The output is a JSON object containing different elements. Among those, the access token is present.

- The access token value is parsed and extracted to then be used in the header for the data request.

- The listview (Salesforce resource) is queried and succesfully unmarshalled. The values are currently printed (not returned).


## Next Steps

- Input for the script to be accepted via enviroment variables. [DONE]
- Parse JSON output and extract only the "access token" value. [DONE]
- Pass the access token value to next section of the script. [DONE]
- Modify the downloaded resource (currently sObject/case) to a resource that provides the list of active cases from the engineer.
- Avoid the usage of external tools (jq and/or grep), build the presentation of data in the source code.
- Allow the user to specify a case ID to get additional information about it.
- Containerize the application.
- Get the attachment from Salesforce:
    - Direct attachments from Salesforce.
    - Attachments from third party integrated tool like S-Drive.







