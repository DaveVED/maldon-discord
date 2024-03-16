package commands

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

// WelcomeNewMember sends a welcome message to new members.
func WelcomeNewMember(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
    welcomeMessage := fmt.Sprintf("Welcome %s to the server!", m.Member.User.Username)
    s.ChannelMessageSend("1218320382717595748", welcomeMessage)
}

