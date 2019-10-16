# Tweet Scheduler

This is a small go program that tweets on a schedule.

Lots and lots of this code was taken from https://github.com/go-study-group/gostudygroup-bot, and that was initially written by [@ankur-anand](https://github.com/ankur-anand). Thank you so much to you, Ankur!

# Config

## Twitter Specific

- `TWITTER_CONSUMER_KEY`=xxxxxx
- `TWITTER_CONSUMER_SECRET`=xxxxxx
- `TWITTER_ACCESS_TOKEN`=xxxxxxxx
- `TWITTER_ACCESS_TOKEN_SECRET`=xxxxxxx

## Application

- `GO_ENV`
  - Should be "production" for production uses otherwise intended tweet will not get twitted
- `TWITTER_POST_API_TOKEN`
  - Will be used to verify every post call to tweet, the body parameter `token` value should be same.
- `PORT`: the port the app webserver should run on
