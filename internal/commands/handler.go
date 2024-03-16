package commands

import (
    "log"
    
    "github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    log.Printf("Message received: '%s'", m.Content)

    if m.Author.ID == s.State.User.ID {
        log.Println("Ignoring bot's own message.")
        return 
    }

    if m.Content == "!ping" {
        PingMessageCreate(s, m)
    }
}

