package handler

import (
    "github.com/bwmarrin/discordgo"
    "log"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Enclose the message content in quotes for clarity
    log.Printf("Message received: '%s'", m.Content)

    // Ignore bot's own messages
    if m.Author.ID == s.State.User.ID {
        log.Println("Ignoring bot's own message.")
        return 
    }

    // Trim whitespace and check if content is exactly "ping"
    if m.Content == "!ping" {
        log.Println("Responding to 'ping' with 'pong'.")
        s.ChannelMessageSend(m.ChannelID, "pong")
    } else {
        log.Println("Message did not match 'ping'.")
    }
}

