package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	matrixbot "github.com/binaryplease/matrix-bot"
	log "github.com/sirupsen/logrus"
)

//GiteaBot is the main struct to hold the bot
type GiteaBot struct {
	*matrixbot.MatrixBot
	//map rooms to tokens
	Tokens map[string]string
	db     *GiteaDB
}

//NewGiteaBot creates a new bot form user credentials
func NewGiteaBot(user, pass string, DBPath string) *GiteaBot {

	bot, err := matrixbot.NewMatrixBot(user, pass, "gitea")

	if err != nil {
		log.Fatal(err)
	}
	db := NewGiteaDB(DBPath)

	gbot := &GiteaBot{
		bot,
		db.GetAll(),
		db,
	}

	bot.RegisterCommand("secret", 0, "Request token for a webhook", gbot.handleCommandSecret)

	return gbot

}

func tokenGenerator() string {
	//TODO make token length configurable
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (gb *GiteaBot) checkToken(room, token string) bool {
	return token == gb.db.GetToken(room)
}

func (gb *GiteaBot) handleCommandSecret(message, room, sender string) {

	//Check if a parameter was supplied (room for which to generate the token)
	msgParts := strings.Split(message, " ")
	if len(msgParts) != 3 {
		gb.SendTextToRoom(room, "!gitea secert expects exactly on parameter, a room for which to request a token.\n Usage: !gitea secret <room id> \n\n e.g. !gitea secert !FoJFjcBoIJyKuPnDFf:matrix.org")
		return
	}

	reqRoom := msgParts[2]

	//Basic check if the roomID is properly formatted. The user might mistake it for the alias or address
	if !strings.HasPrefix(reqRoom, "!") {
		gb.SendTextToRoom(room, "Room IDs start with an exclamation mark\n\n e.g. !gitea secert !FoJFjcBoIJyKuPnDFf:matrix.org \n\n This is *not* the same as the rooms name or alias!")
		return
	}

	//Check if room already has a token, if so display it.
	if gb.Tokens[reqRoom] != "" {
		gb.SendTextToRoom(room, "This room already has a token. Your secert token is:")
		gb.SendTextToRoom(room, gb.Tokens[reqRoom])
		return
	}

	// If everything is fine, generate a token and add it to the db
	token := tokenGenerator()
	gb.Tokens[reqRoom] = token
	gb.db.Set(reqRoom, token)

	//Print the token and help to the room it was requested from
	gb.SendTextToRoom(room, "Your secert token is:")
	gb.SendTextToRoom(room, token)
	gb.SendTextToRoom(room, "Now, set up a weebhook in gitea with that token as secret")

	httpHost := cfg.Section("http").Key("http_host").String()
	httpPort := cfg.Section("http").Key("http_port").String()
	httpURI := cfg.Section("http").Key("http_uri").String()
	gb.SendTextToRoom(room, httpHost+":"+httpPort+httpURI+room)
}
