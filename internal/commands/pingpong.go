package commands

import (
    "github.com/bwmarrin/discordgo"
)

// PingMessageCreate handles messages and responds to "ping" with "pong".
func PingMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    s.ChannelMessageSend(m.ChannelID, "pong")
}
