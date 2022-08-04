# salesforce-backlog-cli (Dev branch)

accessToken.go successfully requests an access token to the Salesforce Instace:

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

❯ ./accessToken
{"access_token":"XYXYXYXYXYYXYXYXYXYXYXXYYXYXYYXYXYYXYXYYXYXYYXYXYXYXYXYYX","instance_url":"https://XXXXXXXX.salesforce.com","id":"https://login.salesforce.com/id/XXXXXXXXXXXX/XXXXXXXXXXXXXX","token_type":"Bearer","issued_at":"555555555555555","signature":"XXXXXXXXXXXXXXXXXXXX"}
```

- The script reads the enviroment variables (EMAIL, PASS, SF, CLID, CLSE, SECK) from the user's terminal and uses them as input.
- The output is a JSON object containing different elements. Among those, the access token is present.

## Next Steps

- Input for the script to be accepted via enviroment variables. [DONE]
- Parse JSON output and extract only the "access token" value.
- Pass the access token value to next section of the script. (To download the backlog)


