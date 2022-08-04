# salesforce-backlog-cli (Dev branch)

accessToken.go successfully requests an access token to the Salesforce Instace:

```
❯ ./main -user=$NAME -pass=$PASS -sf=$SF -clid=$CLID -clse=$CLSE -seck=$SECK
{"access_token":"XYXYXYXYXYYXYXYXYXYXYXXYYXYXYYXYXYYXYXYYXYXYYXYXYXYXYXYYX","instance_url":"https://XXXXXXXX.salesforce.com","id":"https://login.salesforce.com/id/XXXXXXXXXXXX/XXXXXXXXXXXXXX","token_type":"Bearer","issued_at":"555555555555555","signature":"XXXXXXXXXXXXXXXXXXXX"}
```

- The input for the script is passed with flags: user, pass, sf, clid, clse, seck.
- The output is a JSON object containing different elements. Among those, the access token is present.

## Next Steps

- Input for the script to be accepted via enviroment variables. (Seems more suitable)
- Parse JSON output and extract only the "access token" value.
- Pass the access token value to next section of the script. (To download the backlog)


