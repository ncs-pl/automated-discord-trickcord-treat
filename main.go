package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	channel string
	token string
)

func init() {
	flag.StringVar(&channel, "channel", "INVALID", "--channel <CHANNEL ID>")
	flag.StringVar(&token, "token", "INVALID", "--token <USER TOKEN !PRIVATE!>")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}

func ready(s *discordgo.Session, _ *discordgo.Ready) {
	fmt.Println("Automation ready!  Press CTRL-C to exit.")
	c, err := s.Channel(channel)
	if err != nil {
		fmt.Println("Cannot find channel with ID " + channel + "->\n" + err.Error())
	} else {
		fmt.Println("Channel: #" + c.Name + " with ID " + c.ID)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if channel == m.ChannelID {
		if m.Author.ID == "755580145078632508" && m.Embeds[0].Title == "A trick-or-treater has stopped by!" {
			c, _ := s.Channel(m.ChannelID)
			fmt.Println("Trick Or Treat in Channel #" + c.Name)
			if strings.Contains(m.Embeds[0].Description, "h!trick") {
				_,_ = s.ChannelMessageSend(m.ChannelID, "h!trick")
				fmt.Println("Sent `h!trick`")
				seemsLegit(s, m)
			} else if strings.Contains(m.Embeds[0].Description, "h!treat") {
				_, _ = s.ChannelMessageSend(m.ChannelID, "h!treat")
				fmt.Println("Sent `h!treat`")
				seemsLegit(s, m)
			} else {
				fmt.Println("Unknown, did nothing...")
			}
		}
	}
}

func seemsLegit(s *discordgo.Session, m *discordgo.MessageCreate) {
	time.Sleep(1 * time.Second)
	ans := [5]string{"Yeah!", "Boooom", "!!!!!", "fast af", "yes"}
	_, _ = s.ChannelMessageSend(m.ChannelID, ans[rand.Intn(len(ans))])
}
