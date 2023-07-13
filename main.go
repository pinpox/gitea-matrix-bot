package main

import (
	// "fmt"
	"github.com/go-ini/ini"
	"github.com/alecthomas/kingpin/v2"

	"time"
	// "time"
	log "github.com/sirupsen/logrus"
)

var cfg *ini.File
var err error
var mygiteabot *GiteaBot

var (
	verboseFlag  = kingpin.Flag("verbose", "Verbose mode, displays additional information.").Short('v').Default("false").Bool()
	configFlag   = kingpin.Flag("config", "Configuration file to use").Short('c').Default("config.ini").String()
	initDBFlag   = kingpin.Flag("initdb", "Initialize the database. If it exists, it will be overwritten!").Default("false").Bool()
	syncSecsFlag = kingpin.Flag("sync", "Matrix synchronizing interval").Default("1").Int()
)

func init() {

	//Parse flags and set log-level
	kingpin.Version("1.0.0")
	kingpin.Parse()

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
	log.Debugf("Using configuration file %s", *configFlag)
	cfg, err = ini.Load(*configFlag)

	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}

	matrixUser := cfg.Section("matrix").Key("matrix_user").String()
	matrixPass := cfg.Section("matrix").Key("matrix_pass").String()
	matrixHost := cfg.Section("matrix").Key("matrix_pass").String()
	dbPath := cfg.Section("bot").Key("db_path").String()

	//Set up the bot

	// func NewGiteaBot(user, pass, host, string, DBPath string) *GiteaBot {
	mygiteabot = NewGiteaBot(matrixUser, matrixPass, matrixHost, dbPath)

	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Bot created")
}

func main() {

	log.Info("Setting up POST-listener")
	go func() {
		for {
			mygiteabot.Sync()
			time.Sleep(time.Duration(*syncSecsFlag) * time.Second)
		}
	}()
	setupListener()
}
