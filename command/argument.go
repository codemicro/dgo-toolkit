package command

import "github.com/bwmarrin/discordgo"

// Argument is used for parsing arguments in commands.
type Argument struct {
	Name string
	Type ArgumentType
	Default func(session *discordgo.Session, message *discordgo.MessageCreate) (interface{}, error)
}

// ArgumentType defines a possible argument, how to parse it and generates relevant help text
type ArgumentType interface {
	// Parse should consume the argument is has parsed
	Parse(content *string) (interface{}, error)
	Help(name string) string
}
