package commands

import "github.com/bwmarrin/discordgo"

func WhoamiMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "You are "+m.Author.Username)
}
