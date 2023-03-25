package main

import (
	"os"
	"os/signal"
	"syscall"

	"akika.fr/discord/commands"
	"akika.fr/discord/event"
	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Args[1])
	if err != nil {
		println("Can't connect to discord ")
		return
	}

	err = discord.Open()
	if err != nil {
		println("Provided token can't be use")
		return
	}
	event.Message(discord)
	println("Bot started")

	commands.Init(discord)
	commands.Handler(discord)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	commands.Remove(discord)
	discord.Close()
}
