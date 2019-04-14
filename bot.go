package main

import (
	"fmt"
	"github.com/matrix-org/gomatrix"
	"regexp"
	"strings"
)

//MatrixBot struct to hold the bot and it's methods
type MatrixBot struct {
	//Map a repository to matrix rooms
	Client     *gomatrix.Client
	matrixPass string
	matrixUser string
	Handlers   []MatrixBotCommandHandler
}

type MatrixBotCommandHandler struct {
	//The pattern or command to handle
	Pattern string

	//The minimal power requeired to execute this command
	MinPower int

	//The function to handle this command
	Handler func(message, room, sender string)
}

func (gb *MatrixBot) getSenderPower(sender string) int {
	//TODO
	return 100
}

func (gb *MatrixBot) RegisterCommand(pattern string, minpower int, handler func(message string, room string, sender string)) {
	mbch := MatrixBotCommandHandler{
		Pattern:  pattern,
		MinPower: minpower,
		Handler:  handler,
	}
	fmt.Println("Registered command: " + pattern)
	gb.Handlers = append(gb.Handlers, mbch)
}

func (gb *MatrixBot) handleCommands(message, room, sender string) {

	//Don't do anything if the sender is the bot itself
	//TODO edge-case: bot has the same name as a user but on a different server
	if strings.Contains(sender, gb.matrixUser) {
		return
	}

	for _, v := range gb.Handlers {
		r, _ := regexp.Compile(v.Pattern)
		if r.MatchString(message) {
			if v.MinPower <= gb.getSenderPower(sender) {
				v.Handler(message, room, sender)
			} else {
				gb.SendToRoom(room, "You have not enough power to execute this command ("+v.Pattern+"). Your power: "+string(gb.getSenderPower(sender))+", requeired: "+string(v.MinPower))
			}
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

//SendToRoom sends a message to a specified room
func (gb *MatrixBot) SendToRoom(room, message string) {
	_, err = gb.Client.SendText(room, message)
	if err != nil {
		panic(err)
	}
}

//NewGiteaBot creates a new bot form user credentials
func NewMatrixBot(user, pass string) (*MatrixBot, error) {

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

	bot := &MatrixBot{
		matrixPass: pass,
		matrixUser: user,
		Client:     cli,
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
