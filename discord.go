package main

import (
	"fmt"
	"io/ioutil" //reading files
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/emirpasic/gods/sets/hashset"
)

var dg *discordgo.Session
var gamers map[string]*hashset.Set = make(map[string]*hashset.Set)
var minPlayers int = 2

func discord() { // init function
	var err error
	dg, err = discordgo.New("Bot " + key())
	if err != nil {
		panic(err)
	}
	dg.Open()
	fmt.Println("Started and connected")
	dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessageReactions | discordgo.IntentsDirectMessages | discordgo.IntentsGuildMessageReactions
	dg.AddHandler(messageListen) // new messages will be sent here
	dg.AddHandler(emojiAddListen)
	dg.AddHandler(emojiRmListen)
	<-sig
	dg.Close()
}
func key() string {
	token, err := ioutil.ReadFile(".config")
	if err != nil {
		panic("Please place bot token in a file named \".config\" in the install directory")
	}
	return string(token)
}
func messageListen(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot { // don't recieve orders from itself and bots
		return
	}
	if strings.ToLower(m.Message.Content) == "start discord among us" {
		dg.MessageReactionAdd(m.ChannelID, m.Message.ID, "☑️")
		gamers[m.Message.ID] = hashset.New()
		gamers[m.Message.ID].Add(m.Message.Author.ID)
		timer := time.NewTimer(time.Second * 30)
		<-timer.C
		if gamers[m.Message.ID].Size() < minPlayers {
			s.ChannelMessageSend(m.ChannelID, "Not enough players! ("+strconv.Itoa(minPlayers)+" required)")
		} else {
			startGame()
		}
	}
}
func emojiAddListen(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.MessageReaction.UserID == s.State.User.ID { // don't recieve orders from itself
		return
	}
	value, exists := gamers[m.MessageID]
	if exists && m.Emoji.Name == "☑️" {
		value.Add(m.MessageReaction.UserID)
	}
}
func emojiRmListen(s *discordgo.Session, m *discordgo.MessageReactionRemove) {
	if m.MessageReaction.UserID == s.State.User.ID { // don't recieve orders from itself
		return
	}
	value, exists := gamers[m.MessageID]
	if exists && m.Emoji.Name == "☑️" {
		value.Remove(m.MessageReaction.UserID)
	}
}
