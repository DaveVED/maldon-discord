package commands

import (
    "github.com/bwmarrin/discordgo"
)

// PingMessageCreate handles messages and responds to "ping" with "pong".
func PingMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Content == "!ping" {
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
}
