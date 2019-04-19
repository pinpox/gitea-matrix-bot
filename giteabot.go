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

func (gb *GiteaBot) Send(token, message string) {
	for k, v := range gb.Tokens {
		if v == token {
			gb.SendToRoom(k, message)
		}
	}
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
	// bot.RegisterCommand("set", 0, "Set an existing token for the room", gbot.handleCommandSet)
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

// func (gb *GiteaBot) handleCommandSet(message, room, sender string) {

// 	// Get the parameter(s) given to the command
// 	args := strings.Split(message, " ")

// 	// Display help/error if more than one argument is given
// 	if len(args) != 3 {
// 		gb.SendToRoom(room, "set expects exactly one argument")
// 		gb.SendToRoom(room, "!gitea set <token>")
// 	} else {
// 		// Display help/error if the token has the wrong length
// 		if len(args[2]) != 20 {
// 			gb.SendToRoom(room, "Tokens have a length of 20 characters")
// 		} else {
// 			// If the token seems ok, set it for the room
// 			gb.SendToRoom(room, "Setting token for this room to:")
// 			gb.SendToRoom(room, args[2])
// 			gb.Tokens[room] = args[2]
// 			gb.db.Update(gb.Tokens)
// 		}
// 	}
// }

func (gb *GiteaBot) handleCommandSecret(message, room, sender string) {

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
