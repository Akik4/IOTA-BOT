package commands

import "github.com/bwmarrin/discordgo"

func Remove(s *discordgo.Session) {
	for _, v := range RegisteredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			println("Suppresion impossible")
		}
	}
}
