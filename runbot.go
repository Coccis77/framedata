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

var globalGetMove data.GetMove

func RunBot(botToken string, g data.GetMove) {
	dg, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}

	globalGetMove = g

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

	response, err := data.FindData(strings.Replace(message.Content, "!", "", 1), globalGetMove)
	if err != nil {
		log.Print(err.Error())
		return
	}

	if len(response) > 0 {
		session.ChannelMessageSend(message.ChannelID, "```" + response + "```")
	}
}