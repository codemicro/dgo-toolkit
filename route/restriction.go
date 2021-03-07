package route

import "github.com/bwmarrin/discordgo"

// CommandRestriction is a function that returns true if a command can be run based on the current state (eg. user, roles,
// channel, etc)
type CommandRestriction func(session *discordgo.Session, message *discordgo.MessageCreate) (bool, error)

// TODO: premade restriction templates, eg restriction by role, by channel, etc...