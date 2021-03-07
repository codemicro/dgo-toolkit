package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type MessageRunFunc func(ctx *MessageContext) error

type MessageContext struct {
	Session   *discordgo.Session
	Message   *discordgo.MessageCreate
	Arguments map[string]interface{}
}

// onMessageCreate is a callback function to be used with a DiscordGo session that iterates through all registered
// commands and runs the first one that it finds that matches
func (b *Kit) onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// check if the message has a given prefix
	var trimmedContent string
	for _, prefix := range b.Prefixes {
		// slightly modified version of strings.HasPrefix
		if b.hasPrefix(message.Content, prefix) {
			trimmedContent = b.trimPrefix(message.Content, prefix)
		}
	}

	if trimmedContent == "" {
		// no command? nothing for us to do
		return
	}

	// iterate registered commands
	for _, cmd := range b.commandSet {
		if cmd.detectRegexp.MatchString(trimmedContent) {

			// check if all restrictions pass
			ok := true
			if cmd.Restrictions != nil {
				for _, rf := range cmd.Restrictions {
					rfOk, err := rf(session, message)
					if err != nil {
						b.handleError(err)
						return // TODO: could something else be done here?
					}
					ok = ok && (rfOk || b.DebugMode) // use debug mode to ignore restrictions
				}
			}

			if !ok { // if the restrictions have not been met
				err := session.MessageReactionAdd(message.ChannelID, message.ID, "⚠")
				if err != nil {
					b.handleError(err)
				}
				break
			}

			// remove command text
			{
				tcx := cmd.detectRegexp.Split(trimmedContent, -1)
				trimmedContent = strings.TrimSpace(tcx[1])
			}

			// parse arguments
			argumentMap := make(map[string]interface{})
			if cmd.Arguments != nil {
				for _, arg := range cmd.Arguments {

					var val interface{}
					var failMessage string

					if len(trimmedContent) == 0 { // if there's nothing left to parse
						if arg.Default != nil { // if there's a default available

							dv, err := arg.Default(session, message)
							if err != nil {
								b.handleError(err)
								return
							}
							val = dv

						} else {
							failMessage = "argument missing"
						}

					} else { // otherwise parse from the available text
						var err error
						val, err = arg.Type.Parse(&trimmedContent)
						if err != nil {
							failMessage = err.Error()
						}
					}

					if failMessage != "" {
						cont := fmt.Sprintf("❌ `%s`: %s", arg.Name, failMessage)
						_, err := session.ChannelMessageSend(message.ChannelID, cont)
						if err != nil {
							b.handleError(err)
						}
						return
					}

					argumentMap[arg.Name] = val

				}
			}

			ctx := &MessageContext{
				Session:   session,
				Message:   message,
				Arguments: argumentMap,
			}

			err := cmd.Run(ctx)
			if err != nil {
				b.handleError(err, cmd.Name)
			}

			return // no more commands

		}
	}

}
