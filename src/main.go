package main

import (
	"os"
	"os/signal"
	"src/src/commands"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/zekrotja/ken"
	"github.com/zekrotja/ken/store"
)

var err error
var Session *discordgo.Session

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// Getting token from environment variables
	var token = os.Getenv("POCHITA")

	// Initiating new discord bot client
	Session, err = discordgo.New("Bot " + token)
	must(err)
	defer must(Session.Close())

	k, err := ken.New(Session, ken.Options{
		CommandStore: store.NewDefault(),
	})
	must(err)

	// Register commands
	must(k.RegisterCommands(new(commands.TestCommand)))

	defer must(k.Unregister())

	must(Session.Open())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
