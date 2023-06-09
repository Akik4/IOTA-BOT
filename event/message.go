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
			if strings.Split(m.Content, " ")[1] == m.Mentions[0].Mention() {
				ressources.Avatar(s, m, m.Mentions[0])
			} else {
				s.ChannelMessageSend(m.ChannelID, "You can't put text as first argument\nexpected : !avatar <@user>")
			}
		} else {
			ressources.Avatar(s, m, m.Author)
		}
	}
}
