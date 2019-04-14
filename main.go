package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

// import "net/http"

var cfg *ini.File
var err error
var bot *GiteaBot

func init() {
	//Lonad config
	cfg, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	matrixUser := cfg.Section("matrix").Key("matrix_user").String()
	matrixPass := cfg.Section("matrix").Key("matrix_pass").String()
	// botDB := cfg.Section("bot").Key("").String()

	fmt.Println("Creating Bot")
	bot, err = NewGiteaBot(matrixUser, matrixPass)

	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("Setting up POST-listener")
	setupListener()
}
