package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session
var gamers map[string]string = make(map[string]string)

func discord() { // init function

	var err error
	dg, err = discordgo.New("Bot " + key)
	if err != nil {
		panic(err)
	}
	dg.Open()
	fmt.Println("Started and connected")
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessageReactions | discordgo.IntentsDirectMessages | discordgo.IntentsGuildMessageReactions
	go dg.AddHandler(messageListener) // new messages will be sent here
	go dg.AddHandler(reactionListener)
	<-sig
	dg.Close()
}
func messageListener(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot { // don't recieve orders from itself and bots
		return
	}
	if strings.ToLower(m.Message.Content) == "start discord among us" {
		dg.MessageReactionAdd(m.ChannelID, m.MessageReference.MessageID, "ballot_box_with_check")
		gamers[m.Message.ID] = m.Message.Author.ID
	}
}
func reactionListener(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.MessageReaction.UserID == s.State.User.ID { // don't recieve orders from itself and bots
		return
	}
	value, exists := gamers[m.MessageID]
	if exists {
		value = m.MessageReaction.UserID
	}
	fmt.Print(value)
}
