package commands

import (
    "fmt"
    
    "github.com/bwmarrin/discordgo"
)

// WelcomeNewMember sends a welcome message to new members.
func WelcomeNewMember(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
    welcomeChannelID := "1218320382717595748"
    welcomeMessage := fmt.Sprintf("Welcome %s to the server!", m.User.Mention())
    s.ChannelMessageSend(welcomeChannelID, welcomeMessage)
}

