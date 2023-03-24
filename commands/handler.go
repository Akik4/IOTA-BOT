package commands

import "github.com/bwmarrin/discordgo"

var RegisteredCommands map[int]*discordgo.ApplicationCommand

var (
	CommandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"avatar": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
		},
		"userinfo": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

					var img = []*discordgo.MessageEmbedThumbnail{{URL: users.AvatarURL("")}}
					var footer = []*discordgo.MessageEmbedFooter{{Text: "Provided by " + i.Member.User.Username}}
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Embeds: []*discordgo.MessageEmbed{{
								Title: users.Username + "'s information",
								Fields: []*discordgo.MessageEmbedField{
									{
										Name:   "ID",
										Value:  users.ID,
										Inline: true,
									},
									{
										Name:   "TAG",
										Value:  users.Discriminator,
										Inline: true,
									},
									{
										Name:   "JOIN",
										Value:  i.Member.JoinedAt.String(),
										Inline: false,
									},
								},
								Thumbnail: img[0],
								Footer:    footer[0],
								Color:     15418782,
							}},
						},
					})
				}
			} else {
				var img = []*discordgo.MessageEmbedThumbnail{{URL: i.Member.AvatarURL("")}}
				var footer = []*discordgo.MessageEmbedFooter{{Text: "Provided by " + i.Member.User.Username}}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{{
							Title: i.Member.User.Username + "'s avatar",
							Fields: []*discordgo.MessageEmbedField{
								{
									Name:   "ID",
									Value:  i.Member.User.ID,
									Inline: true,
								},
								{
									Name:   "TAG",
									Value:  i.Member.User.Discriminator,
									Inline: true,
								},
								{
									Name:   "JOIN AT",
									Value:  i.Member.JoinedAt.String(),
									Inline: false,
								},
							},
							Thumbnail: img[0],
							Footer:    footer[0],
							Color:     15418782}},
					},
				})
			}
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
