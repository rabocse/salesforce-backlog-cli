# salesforce-backlog-cli (Dev branch)



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
====== FROM BELOW IS A WORK IN PROGRESS ==========

######## CASE ########

Case Number: 1234567
Client's Name: Aleks Esc
Subject's Case: [EMEA] Monitoring Failing
Severity: Sev4 (Low)
Status: Closed
Environment: Company Fake

######## CASE ########

Case Number: 8910111
Client's Name: Aless Sandro
Subject's Case: [EMEA] Possible Memory Leak
Severity: Sev3 (Normal)
Status: Closed
Environment: Italy Corpo Fake

######## CASE ########

Case Number: 1213141
Client's Name: Alex Rabocse
Subject's Case: [AMER] Questions About Kubernetes
Severity: Sev4 (Low)
Status: Solved
Environment: Rabocse Coporate
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







