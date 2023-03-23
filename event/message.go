package event

import (
	"strings"

	"akika.fr/discord/ressources"
	"github.com/bwmarrin/discordgo"
)

func Message(s *discordgo.Session) {
	s.AddHandler(listener)
}

func listener(s *discordgo.Session, m *discordgo.MessageCreate) {

	if s.State.User.ID == m.Author.ID {
		return
	}

	if strings.Split(m.Content, " ")[0] == "!avatar" {
		if len(m.Mentions) > 0 {
			ressources.Avatar(s, m, m.Mentions[0])
		} else {
			ressources.Avatar(s, m, m.Author)
		}
	}
}
