package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"

	_ "github.com/matrix-org/gomatrix"
)

// import "net/http"

//TODO get this vals from config file
// var (
// flagPort is the open port the application listens on
// flagPort   = flag.String("port", "9000", "Port to listen on")
// matrixPass = "oSaiNahqu5ahF5jieBe2UKok"
// matrixUser = "reminder-bot"
// matrixHost = "http://matrix.org"

// postURI = "/post"
// )

var cfg *ini.File
var err error

func init() {

	cfg, err = ini.Load("config.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

}

func main() {

	//TODO Check if already logged in
	//TODO Check if In room
	//TODO Send

	// fmt.Println("Logging in")

	// cli, _ := gomatrix.NewClient("http://matrix.org", "", "")
	// resp, err := cli.Login(&gomatrix.ReqLogin{
	// 	Type:     "m.login.password",
	// 	User:     matrixUser,
	// 	Password: matrixPass,
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// cli.SetCredentials(resp.UserID, resp.AccessToken)

	// fmt.Println("Joining Room")

	// roomID := ""

	// if resp, err := cli.JoinRoom("#test-reminder-bot:matrix.org", "", nil); err != nil {
	// 	panic(err)
	// } else {
	// 	roomID = resp.RoomID
	// }
	// fmt.Println("Sending a message")
	// _, err = cli.SendText(roomID, "test message 0")
	// if err != nil {
	// 	panic(err)

	// }

	fmt.Println("Setting up POST-listener")
	setupListener()

}
