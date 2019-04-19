package main

import (
	"crypto/rand"
	"fmt"
	// "strings"

	matrixbot "github.com/binaryplease/matrix-bot"
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
		panic(err)
	}
	db := NewGiteaDB(DBPath)

	gbot := &GiteaBot{
		bot,
		db.GetAll(),
		db,
	}

	bot.RegisterCommand("secret", 0, "Request token for a webhook", gbot.handleCommandSecret)
	bot.RegisterCommand("reset", 0, "Delete the room's token", gbot.handleCommandReset)

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

func (gb *GiteaBot) handleCommandReset(message, room, sender string) {
	gb.db.Unset(room, gb.Tokens[room])

	_, ok := gb.Tokens[room]
	if ok {
		delete(gb.Tokens, room)
	}
}

func (gb *GiteaBot) handleCommandSecret(message, room, sender string) {

	//TODO make the room a parameter (dont use the current room as room)

	//Check if room already has a token
	if gb.Tokens[room] != "" {
		gb.SendToRoom(room, "This room already has a token. Your secert token is:")
		gb.SendToRoom(room, gb.Tokens[room])
		gb.SendToRoom(room, "To remove all tokens of this room use !gitea reset")
		return
	}

	token := tokenGenerator()
	gb.Tokens[room] = token
	gb.db.Set(room, token)

	gb.SendToRoom(room, "Your secert token is:")
	gb.SendToRoom(room, token)

	gb.SendToRoom(room, "Set up a weebhook in gitea with that token as secret")

	//TODO change this to a real URL or make it configurable
	gb.SendToRoom(room, "http://192.168.2.33:9000/post/"+room)
}
