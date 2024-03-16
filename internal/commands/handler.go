package commands

import (
    "log"
    
    "github.com/bwmarrin/discordgo"
)

// MessageCreate is called whenever a new message is created in a channel.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages sent by the bot itself
    if m.Author.ID == s.State.User.ID {
        return 
    }

    switch m.Content {
    case "!ping":
        PingMessageCreate(s, m)
    case "!whoami":
        WhoamiMessageCreate(s, m)
    default:
        log.Println("Command not recognized:", m.Content)
    }
}
