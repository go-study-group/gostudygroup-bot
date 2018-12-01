## Post a 5 min Notification Tweet to Twitter.

POST /api/v1/tweets/startinfive

### Parameters

POST body

- token (required) - A **"string"** token for every request, that was passed as TWITTER_POST_API_KEY as env variable while running the app.

### Returns

Success

```json
{
  "result": "The tweet has been successfully posted, to Handle [<twitterhandle>] and status ID is [<statusID>]"
}
```

Duplicate error

```json
{
  "result": "Get https://api.twitter.com/1.1/statuses/update.json returned status 403, {\"errors\":[{\"code\":187,\"message\":\"Status is a duplicate.\"}]}"
}
```
