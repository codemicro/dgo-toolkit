package route

import "github.com/bwmarrin/discordgo"

type MessageContext struct {
	Session   *discordgo.Session
	Message   *discordgo.MessageCreate
	Arguments map[string]interface{}
	kit       *Kit
}

func (m *MessageContext) DefaultAllowedMentions() *discordgo.MessageAllowedMentions {
	// This copy is intentional
	n := m.kit.DefaultAllowedMentions
	return &n
}

func (m *MessageContext) SendMessageString(channelId string, content string) (*discordgo.Message, error) {

	return m.Session.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
		Content:         content,
		AllowedMentions: m.DefaultAllowedMentions(),
	})

}

func (m *MessageContext) SendMessageEmbed(channelId string, embed *discordgo.MessageEmbed) (*discordgo.Message, error) {

	return m.Session.ChannelMessageSendComplex(channelId, &discordgo.MessageSend{
		Embed: embed,
		AllowedMentions: m.DefaultAllowedMentions(),
	})

}

