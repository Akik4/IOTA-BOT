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
				var Gender discordgo.ApplicationCommandOptionType
				var req bool
				t := V.Settings.Type[i]
				if t == "user" {
					Gender = discordgo.ApplicationCommandOptionUser
				} else if t == "string" {
					Gender = discordgo.ApplicationCommandOptionString
				} else if t == "attachment" {
					Gender = discordgo.ApplicationCommandOptionAttachment
				} else if t == "boolean" {
					Gender = discordgo.ApplicationCommandOptionBoolean
				} else if t == "channel" {
					Gender = discordgo.ApplicationCommandOptionChannel
				} else if t == "int" {
					Gender = discordgo.ApplicationCommandOptionInteger
				} else if t == "mentionable" {
					Gender = discordgo.ApplicationCommandOptionMentionable
				} else if t == "number" {
					Gender = discordgo.ApplicationCommandOptionNumber
				} else if t == "role" {
					Gender = discordgo.ApplicationCommandOptionRole
				} else if t == "subcommand" {
					Gender = discordgo.ApplicationCommandOptionSubCommand
				} else if t == "subcommandgroup" {
					Gender = discordgo.ApplicationCommandOptionSubCommandGroup
				}

				if V.Settings.Required[i] == "true" {
					req = true
				} else {
					req = false
				}
				Opt = append(Opt, &discordgo.ApplicationCommandOption{
					Name:        V.Settings.Name[i],
					Description: V.Settings.Description[i],
					Type:        Gender,
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
