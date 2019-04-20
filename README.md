[![Build Status](https://travis-ci.org/binaryplease/gitea-matrix-bot.svg?branch=master)](https://travis-ci.org/binaryplease/gitea-matrix-bot)
[![GoDoc](https://godoc.org/github.com/binaryplease/gitea-matrix-bot?status.svg)](https://godoc.org/github.com/binaryplease/gitea-matrix-bot)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/gitea-matrix-bot)](https://goreportcard.com/report/github.com/binaryplease/gitea-matrix-bot)
[![codecov](https://codecov.io/gh/binaryplease/gitea-matrix-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/gitea-matrix-bot)



# gitea-matrix-bot (WORK IN PROGRESS!)

A bot to listen for [gitea](https://gitea.io) webhooks and post to a [matrix](https://matrix.org) channel

# Usage

## Configuration
Copy config.ini.example to config.ini or creat a file called config.ini with the
following options:

```ini
[http]
# The path the listener will expect the post data
http_uri = "/post"
# The port the listener will listen on
http_port = "9000"


[matrix]
# The matrix server to connect to
matrix_host = "http://matrix.org"
# The matrix room to post to
matrix_room = "#my-awesome-room:matrix.org"
# User credentions of the bot for posting to the room
matrix_pass = "supersecretpass"
matrix_user = "my-awesome-bot"


[bot]
# Path of the database to be used
db_path = "./tokens.db"
```

Then start the bot. It will listen on the configured URI for incoming gitea
hooks.

## Create gitea hook

Create a new webhook in gitea. You an add a webhook to a single repository or
add a default webhook that will apply to all you repos (recommended)

- Choose the Webhook type `Gitea` for you Webhook
- Configure which events you want to send. You can also select `All Events`
- Set the Target URL to your host + the value you configured in the `config.ini` file. (e.g. `http://myserver:9000/post`)
- Make sure that port is reachable/forwarded if you are behind a router and
	check your firewall
- Set `POST Content Type` to `applicatino/json`
- Add a secret (passphrase, password)

It will probably look like this:
![gitea scrot](./gitea-scrot.png "Gitea Screenshot")
