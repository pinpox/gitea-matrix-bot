package main

import (
	"crypto/rand"
	"fmt"
	"strings"

	matrixbot "github.com/binaryplease/matrix-bot"
)

//GiteaBot is the main struct to hold the bot
type GiteaBot struct {
	*matrixbot.MatrixBot
	//map rooms to tokens
	Tokens map[string]string
}

func (gb *GiteaBot) Send(token, message string) {
	for k, v := range gb.Tokens {
		if v == token {
			gb.SendToRoom(k, message)
		}
	}
}

// func contains(s []string, e string) bool {
//	for _, a := range s {
//		if a == e {
//			return true
//		}
//	}
//	return false
// }

//NewGiteaBot creates a new bot form user credentials
func NewGiteaBot(user, pass string) *GiteaBot {

	tokens := make(map[string]string)

	bot, err := matrixbot.NewMatrixBot(user, pass)

	if err != nil {
		panic(err)
	}

	gbot := &GiteaBot{
		bot,
		tokens,
	}

	bot.RegisterCommand("!gitea help", 0, gbot.handleCommandHelp)
	bot.RegisterCommand("!gitea secret", 0, gbot.handleCommandSecret)
	bot.RegisterCommand("!gitea set", 0, gbot.handleCommandSet)

	return gbot

}

func tokenGenerator() string {
	//TODO make token length configurable
	b := make([]byte, 20)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (gb *GiteaBot) handleCommandSet(message, room, sender string) {

	// Get the parameter(s) given to the command
	args := strings.Split(message, " ")

	// Display help/error if more than one argument is given
	if len(args) != 3 {
		gb.SendToRoom(room, "set expects exactly one argument")
		gb.SendToRoom(room, "!gitea set <token>")
	} else {
		// Display help/error if the token has the wrong length
		if len(args[2]) != 20 {
			gb.SendToRoom(room, "Tokens have a length of 20 characters")
		} else {
			// If the token seems ok, set it for the room
			gb.SendToRoom(room, "Setting token for this room to:")
			gb.SendToRoom(room, args[2])
			gb.Tokens[room] = args[2]
		}
	}
}

func (gb *GiteaBot) handleCommandSecret(message, room, sender string) {

	//Check if room already has a token
	if gb.Tokens[room] != "" {
		gb.SendToRoom(room, "This room already has a token. Your secert token is:")
		gb.SendToRoom(room, gb.Tokens[room])
		return
	}

	token := tokenGenerator()
	gb.Tokens[room] = token

	gb.SendToRoom(room, "Your secert token is:")
	gb.SendToRoom(room, token)

	gb.SendToRoom(room, "Set up a weebhook in gitea with that token as secret")

	//TODO change this to a real URL or make it configurable
	gb.SendToRoom(room, "http://192.168.2.33:9000/post/")
}

func (gb *GiteaBot) handleCommandHelp(message, room, sender string) {
	//TODO maybe make this help auto-generated for bots in general?
	helpMsg := `

I'm your friendly Gitea Bot!

You can invite me to any matrix room to get updates on subscribed gitea repositorys.
The following commands are avaitible:

!sub user/repo       Subscribe to a repository
!unsub user/repo     Remove subscription to a repository
!listsubs            List the room's subscriptions
!help                Display this message

Some of the commands might require admin powers!

`
	gb.SendToRoom(room, helpMsg)
}
