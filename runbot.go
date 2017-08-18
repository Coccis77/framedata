package framedata

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os/signal"
	"syscall"
	"os"
	"strings"
	"github.com/fmicaelli/framedata/data"
)

// TODO Find a way to make a better injection
var globalMove data.Move

func RunBot(botToken string, m data.Move) {
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	globalMove = m

	dg.AddHandler(onMessage)

	dg.Open()
	defer dg.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(message.Content, "!") {
		return
	}

	response, err := data.FindData(strings.Replace(message.Content, "!", "", 1), globalMove)
	if err != nil {
		log.Print(err.Error())
		return
	}

	if len(response) > 0 {
		session.ChannelMessageSend(message.ChannelID, "```" + response + "```")
	}
}