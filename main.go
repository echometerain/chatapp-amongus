package main

import (
	"fmt"
	"os"
	"os/signal" //good practice for handling stop signals
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func main() {
	var err error
	dg, err = discordgo.New("Bot " + key)
	if err != nil {
		panic(err)
	}
	dg.Open()
	fmt.Println("Started and connected")
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(listener) //new messages will be sent to listener()

	sig := make(chan os.Signal, 1) //shuts down the program nicely
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sig
	dg.Close()
}
func listener(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot { //can't recieve orders from youself and other bots
		return
	}
}
