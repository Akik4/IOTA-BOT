package handler

import "github.com/bwmarrin/discordgo"

func AvatarHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	if len(options) > 0 {
		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		if option, ok := optionMap["user"]; ok {

			var users *discordgo.User
			var err error

			users, err = s.User(option.UserValue(nil).ID)
			if err != nil {
				println("erreur :", err)
			}

			var img = []*discordgo.MessageEmbedImage{{URL: users.AvatarURL("")}}
			var footer = []*discordgo.MessageEmbedFooter{{Text: "Provided by " + i.Member.User.Username}}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{{Title: users.Username + "'s avatar", Image: img[0], Footer: footer[0], Color: 15418782}},
				},
			})
		}
	} else {
		var img = []*discordgo.MessageEmbedImage{{URL: i.Member.AvatarURL("")}}
		var footer = []*discordgo.MessageEmbedFooter{{Text: "Provided by " + i.Member.User.Username}}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{{Title: i.Member.User.Username + "'s avatar", Image: img[0], Footer: footer[0], Color: 15418782}},
			},
		})
	}
}
