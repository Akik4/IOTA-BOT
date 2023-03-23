package event

import (
	"github.com/bwmarrin/discordgo"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "avatar",
			Description: "Show users avatar",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "User",
					Type:        discordgo.ApplicationCommandOptionUser,
				},
			},
		},
	}
	commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"avatar": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			if len(options) > 0 {
				optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

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
							Embeds: []*discordgo.MessageEmbed{{Title: users.Username + "'s avatar", Image: img[0], Footer: footer[0]}},
						},
					})
				}
			} else {
				var img = []*discordgo.MessageEmbedImage{{URL: i.Member.AvatarURL("")}}
				var footer = []*discordgo.MessageEmbedFooter{{Text: "Provided by " + i.Member.User.Username}}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Embeds: []*discordgo.MessageEmbed{{Title: i.Member.User.Username + "'s avatar", Image: img[0], Footer: footer[0]}},
					},
				})
			}
		},
	}
)

var registeredCommands map[int]*discordgo.ApplicationCommand

func Commands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			println("Error :", err.Error())
		}
		registeredCommands[i] = cmd
	}
}

func RemoveCommands(s *discordgo.Session) {
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			println("Suppresion impossible")
		}
	}
}
