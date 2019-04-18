package main

import (
	"fmt"
	"github.com/binaryplease/matrix-bot"
)

//GiteaBot is the main struct to hold the bot
type GiteaBot struct {
	*matrixbot.MatrixBot
	Subscriptions map[string][]string
}

//SendMessageToRooms sends a message to all roomes that have subscribed to the repo
func (gb *GiteaBot) SendMessageToRooms(repo, message string) {
	for _, v := range gb.Subscriptions[repo] {
		_, err = gb.Client.SendText(v, message)
		if err != nil {
			panic(err)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//NewGiteaBot creates a new bot form user credentials
func NewGiteaBot(user, pass string) *GiteaBot {

	subs := make(map[string][]string)

	bot, err := matrixbot.NewMatrixBot(user, pass)

	if err != nil {
		panic(err)
	}

	gbot := &GiteaBot{
		bot,
		subs,
	}

	bot.RegisterCommand("!sub_user", 0, gbot.handleCommandAddSubRepo)
	bot.RegisterCommand("!sub_repo", 0, gbot.handleCommandAddSubUser)

	bot.RegisterCommand("!unsub_repo", 0, gbot.handleCommandRemoveSubRepo)
	bot.RegisterCommand("!unsub_user", 0, gbot.handleCommandRemoveSubUser)

	bot.RegisterCommand("!listsubs", 0, gbot.handleCommandListSubs)
	bot.RegisterCommand("!help", 0, gbot.handleCommandHelp)

	return gbot

}

func (gb *GiteaBot) handleCommandListSubs(message, room, sender string) {

	repos := ""

	for k, repo := range gb.Subscriptions {
		for _, subscriber := range repo {
			if subscriber == room {
				repos = repos + "\n -" + k
			}
		}
	}

	if repos == "" {
		gb.SendToRoom(room, "This room has not subscribed to any repositorys.")
	} else {
		msg := "This room has is subscribed to the following repositorys:" + repos
		gb.SendToRoom(room, msg)
	}
}

func (gb *GiteaBot) handleCommandAddSub(message string, room, sender string) {

	repo := message[5:]
	if !contains(gb.Subscriptions[repo], room) {
		gb.Subscriptions[repo] = append(gb.Subscriptions[repo], room)
		gb.SendToRoom(room, "Subscribed to: "+repo)
	} else {
		gb.SendToRoom(room, "This room has already subscribed to: "+repo)
	}
}

func (gb *GiteaBot) handleCommandRemoveSub(message, room, sender string) {

	repo := message[7:]
	fmt.Println(gb.Subscriptions[repo])

	if contains(gb.Subscriptions[repo], room) {

		var tmp []string

		for _, v := range gb.Subscriptions[repo] {
			if v != room {
				fmt.Println("readding '" + v + "'" + "because it is not equal to '" + room + "'")
				tmp = append(tmp, v)
			} else {
				gb.SendToRoom(room, "Un-subscribed from: "+repo)
			}
		}
		gb.Subscriptions[repo] = tmp
	} else {
		gb.SendToRoom(room, "This room has not subscribed to: "+repo)
	}

}

func (gb *GiteaBot) handleCommandHelp(message, room, sender string) {
	helpMsg := `

I'm your friendly Gitea Bot!

You can invite me to any matrix room to get updates on subscribed gitea repositorys.
The following commands are avaitible:

!sub_repo user/repo       Subscribe to a repository
!sub_user user/repo       Subscribe to a user (all repos)
!unsub_repo user/repo     Remove subscription to a repository
!unsub_user user/repo     Remove subscription to a user
!listsubs                 List the room's subscriptions
!help                     Display this message

Some of the commands might require admin powers!

`
	gb.SendToRoom(room, helpMsg)
}
