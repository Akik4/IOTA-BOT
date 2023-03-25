package commands

import (
	"akika.fr/discord/commands/handler"
	"github.com/bwmarrin/discordgo"
)

var RegisteredCommands map[int]*discordgo.ApplicationCommand

var (
	CommandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"avatar": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			handler.AvatarHandler(s, i)
		},
		"userinfo": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			handler.Userinfo(s, i)
		},
	}
)

func Handler(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := CommandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

}
