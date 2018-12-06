## TWITTER Specific

`TWITTER_CONSUMER_KEY`=xxxxxx

`TWITTER_CONSUMER_SECRET`=xxxxxx

`TWITTER_ACCESS_TOKEN`=xxxxxxxx

`TWITTER_ACCESS_TOKEN_SECRET`=xxxxxxx

## GITHUB Specific

`GITHUB_WEBHOOK_REPOAGENDA_SECRET_KEY`=xxxxxxxxxxx

NOTE: `GITHUB_WEBHOOK_REPOAGENDA_SECRET_KEY` webhookSecretKey configured during the webhook trigger for issue.

`GITHUB_ISSUELABELER_INSTALLATION_ID`=xxxxxxxx

NOTE: `GITHUB_ISSUELABELER_INSTALLATION_ID` is the GITHUB APP installation ID for the repo.

`GITHUB_ISSUELABELLER_INTEGERATION_ID`=xxxxxxx

NOTE: `GITHUB_ISSUELABELLER_INTEGERATION_ID` is the GITHUB APP ID.

`GITHUB_ISSUELABELLER_PEMFILE_PATH`=xxxxxxx.pem

NOTE: `GITHUB_ISSUELABELLER_PEMFILE_PATH` complete file location of the pem file path

## APPLICATION Specific

`GO_ENV`=production

NOTE: `GO_ENV` should be "production" for production uses otherwise intended tweet will not get twitted

`TWITTER_POST_API_TOKEN`=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

NOTE: `TWITTER_POST_API_TOKEN` will be used to verify every post call to tweet, the body parameter `token` value should be same.

`PORT`=xxxx // Port for application to run ex - 8080
