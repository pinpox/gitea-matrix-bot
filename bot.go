package main

import (
	"fmt"
	"github.com/matrix-org/gomatrix"
	"regexp"
	"strings"
)

//GiteaBot struct to hold the bot and it's methods
type GiteaBot struct {
	//Map a repository to matrix rooms
	Subscriptions map[string][]string
	Client        *gomatrix.Client
	matrixPass    string
	matrixUser    string
}

func (gb *GiteaBot) handleCommands(message, room, sender string) {

	//Don't do anything if the sender is the bot itself
	//TODO edge-case: bot has the same name as a user but on a different server
	if strings.Contains(sender, gb.matrixUser) {
		return
	}

	//Admin commands
	if gb.isAdmin(sender) {

		//Add a subsription
		rAddSub, _ := regexp.Compile("!sub")
		if rAddSub.MatchString(message) {
			repo := message[4:]
			fmt.Println("You want " + room + " to subscribe to" + repo)
			gb.handleCommandAddSub(room, repo)
			return
		}

		//Remove a subscription
		rRemoveSub, _ := regexp.Compile("!unsub .*")
		if rRemoveSub.MatchString(message) {
			repo := message[6:]
			fmt.Println("You want " + room + " to un-subscribe from " + repo)
			gb.handleCommandRemoveSub(room, repo)
			return
		}

		//List room subs
		rListSubs, _ := regexp.Compile("!listsubs")
		if rListSubs.MatchString(message) {
			fmt.Println("You want to list subs")
			gb.handleCommandListSubs(room)
			return
		}
	}

	//Non-Admin commands

	//Get help
	rHelp, _ := regexp.Compile("!help")
	if rHelp.MatchString(message) {
		fmt.Println("You want to get help")
		gb.handleCommandHelp(room)
		return
	}
}

func (gb GiteaBot) isAdmin(user string) bool {
	//TODO check if admin
	//Check for power in room
	return true
}

func (gb *GiteaBot) handleCommandListSubs(room string) {

	if len(gb.Subscriptions) == 0 {
		gb.SendToRoom(room, "This room has not subscribed to any repositorys.")
		return
	}

	msg := "This room has is subscribed to the following repositorys:"
	for k, repo := range gb.Subscriptions {
		for _, subscriber := range repo {
			if subscriber == room {
				msg = msg + "\n -" + k
			}
		}
	}
	gb.SendToRoom(room, msg)
}

func (gb *GiteaBot) handleCommandAddSub(room, repo string) {
	gb.SendToRoom(room, "Not implemented yet")
}

func (gb *GiteaBot) handleCommandMakeAdmin(room, repo string) {
	gb.SendToRoom(room, "Not implemented yet")
}

func (gb *GiteaBot) handleCommandRemoveSub(room, repo string) {
	gb.SendToRoom(room, "Not implemented yet")
}

func (gb *GiteaBot) handleCommandHelp(room string) {
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

func (gb *GiteaBot) login() {

}

//SendToRoom sends a message to a specified room
func (gb *GiteaBot) SendToRoom(room, message string) {
	_, err = gb.Client.SendText(room, message)
	if err != nil {
		panic(err)
	}
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

//NewGiteaBot creates a new bot form user credentials
func NewGiteaBot(user, pass string) (*GiteaBot, error) {

	fmt.Println("Logging in")

	cli, _ := gomatrix.NewClient("http://matrix.org", "", "")

	resp, err := cli.Login(&gomatrix.ReqLogin{
		Type:     "m.login.password",
		User:     user,
		Password: pass,
	})

	if err != nil {
		return nil, err
	}

	cli.SetCredentials(resp.UserID, resp.AccessToken)

	bot := &GiteaBot{
		matrixPass:    pass,
		matrixUser:    user,
		Subscriptions: make(map[string][]string),
		Client:        cli,
	}

	//Setup Syncer and to handle events
	syncer := cli.Syncer.(*gomatrix.DefaultSyncer)

	//Handle messages send to the channel
	syncer.OnEventType("m.room.message", func(ev *gomatrix.Event) {
		// fmt.Println("\nMessage: ", ev)
		fmt.Println(ev.Sender + " said: \"" + ev.Content["body"].(string) + "\" in room : " + ev.RoomID)
		bot.handleCommands(ev.Content["body"].(string), ev.RoomID, ev.Sender)

	})

	//Handle member events (kick, invite)
	syncer.OnEventType("m.room.member", func(ev *gomatrix.Event) {
		fmt.Println(ev.Sender + " invited bot to " + ev.RoomID)

		if ev.Content["membership"] == "invite" {

			fmt.Println("Joining Room")

			if resp, err := cli.JoinRoom(ev.RoomID, "", nil); err != nil {
				panic(err)
			} else {
				fmt.Println(resp.RoomID)
			}
		}
	})

	//Spawn goroutine to keep checking for events
	go func() {
		for {
			if err := cli.Sync(); err != nil {
				fmt.Println("Sync() returned ", err)
			}
			// Optional: Wait a period of time before trying to sync again.
		}
	}()

	return bot, nil
}
