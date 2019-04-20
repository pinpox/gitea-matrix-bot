package main

import (
	// "fmt"
	"flag"
	"github.com/go-ini/ini"
	// "time"
	log "github.com/sirupsen/logrus"
)

var cfg *ini.File
var err error
var mygiteabot *GiteaBot

var verboseFlag = flag.Bool("v", false, "Display additional information")

func init() {

	flag.Parse()

	if !*verboseFlag {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//Load config
	cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	matrixUser := cfg.Section("matrix").Key("matrix_user").String()
	matrixPass := cfg.Section("matrix").Key("matrix_pass").String()

	mygiteabot = NewGiteaBot(matrixUser, matrixPass, "./tokens.db")

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	log.Info("Starting POST-listener")
	go func() {
		for {
			mygiteabot.Sync()
			// Optional: Wait a period of time before trying to sync again.
		}
	}()
	setupListener()
}
