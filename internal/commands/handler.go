package commands

import (
    "log"
    
    "github.com/bwmarrin/discordgo"
)

// MessageCreate is called whenever a new message is created in a channel.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        log.Println("Ignoring bot's own message.")
        return 
    }

    if m.Content == "!ping" {
        PingMessageCreate(s, m)
    }

    if m.Content == "!whoami" {
        WhoamiMessageCreate(s, m)
    }
}
