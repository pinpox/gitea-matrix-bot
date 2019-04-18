package main

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var cfg *ini.File
var err error
var mygiteabot *GiteaBot

func init() {

	//Load config
	cfg, err = ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	matrixUser := cfg.Section("matrix").Key("matrix_user").String()
	matrixPass := cfg.Section("matrix").Key("matrix_pass").String()

	mygiteabot = NewGiteaBot(matrixUser, matrixPass)

	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("Setting up POST-listener")
	go func() {
		for {
			mygiteabot.Sync()
			// Optional: Wait a period of time before trying to sync again.
		}
	}()
	setupListener()
}
