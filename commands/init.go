package commands

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/bwmarrin/discordgo"
)

var V command

type (
	command struct {
		Name        string
		Description string
		Settings    option
	}
	option struct {
		Load        bool
		Type        []string
		Name        []string
		Description []string
		Required    []string
	}
)

var Commands []*discordgo.ApplicationCommand

func init() {
	folder, err := ioutil.ReadDir("./ressources/commands/")
	if err != nil {
		println("Can't read folder")
	}

	for _, f := range folder {
		file, err := ioutil.ReadFile("./ressources/commands/" + f.Name())
		if err != nil {
			println("error folder ", err.Error())
		}

		err = toml.Unmarshal(file, &V)
		if err != nil {
			println("error unmarshal ", err.Error())
		}

		if V.Settings.Load {

			var Opt []*discordgo.ApplicationCommandOption

			for i := range V.Settings.Type {
				var gender discordgo.ApplicationCommandOptionType
				var req bool
				t := V.Settings.Type[i]

				switch t {
				case "user":
					gender = discordgo.ApplicationCommandOptionUser
				case "string":
					gender = discordgo.ApplicationCommandOptionString
				case "attachment":
					gender = discordgo.ApplicationCommandOptionAttachment
				case "boolean":
					gender = discordgo.ApplicationCommandOptionBoolean
				case "channel":
					gender = discordgo.ApplicationCommandOptionChannel
				case "int":
					gender = discordgo.ApplicationCommandOptionInteger
				case "mentionable":
					gender = discordgo.ApplicationCommandOptionMentionable
				case "number":
					gender = discordgo.ApplicationCommandOptionNumber
				case "role":
					gender = discordgo.ApplicationCommandOptionRole
				case "subcommand":
					gender = discordgo.ApplicationCommandOptionSubCommand
				case "subcommandgroup":
					gender = discordgo.ApplicationCommandOptionSubCommandGroup

				}

				switch V.Settings.Required[i] {
				case "true":
					req = true
				case "false":
					req = false
				}
				Opt = append(Opt, &discordgo.ApplicationCommandOption{
					Name:        V.Settings.Name[i],
					Description: V.Settings.Description[i],
					Type:        gender,
					Required:    req,
				})

			}

			Commands = append(Commands, &discordgo.ApplicationCommand{
				Name:        V.Name,
				Description: V.Description,
				Options:     Opt,
			})
		} else {
			Commands = append(Commands, &discordgo.ApplicationCommand{
				Name:        V.Name,
				Description: V.Description,
			})
		}

	}

}

func Init(s *discordgo.Session) {

	RegisteredCommands := make([]*discordgo.ApplicationCommand, len(Commands))
	for i, v := range Commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			println("Error :", err.Error())
		}
		RegisteredCommands[i] = cmd
	}
}
