package ressources

import "github.com/bwmarrin/discordgo"

func Avatar(s *discordgo.Session, msg *discordgo.MessageCreate, user *discordgo.User) {
	var (
		img = []*discordgo.MessageEmbedImage{
			{
				URL: user.AvatarURL(""),
			},
		}
		footer = []*discordgo.MessageEmbedFooter{
			{
				Text: "Provided by " + msg.Author.Username,
			},
		}
		embed = []*discordgo.MessageEmbed{
			{
				Title:  user.Username + "'s avatar",
				Image:  img[0],
				Footer: footer[0],
				Color:  15418782,
			},
		}
	)
	s.ChannelMessageSendEmbed(msg.ChannelID, embed[0])
}
