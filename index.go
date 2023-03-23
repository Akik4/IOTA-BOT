package main

import (
	"os"
	"os/signal"
	"syscall"

	"akika.fr/discord/event"
	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Args[1])
	if err != nil {
		println("erreur : ", err)
	}

	err = discord.Open()
	if err != nil {
		println("erreur :", err)
	}
	event.Message(discord)
	event.Commands(discord)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	event.RemoveCommands(discord)
	discord.Close()
}
