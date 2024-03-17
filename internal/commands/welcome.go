package commands

import (
    "fmt"

    "github.com/bwmarrin/discordgo"
)

// WelcomeNewMember sends a welcome message to new members.
func WelcomeNewMember(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
    welcomeChannelID := "1218320382717595748"
    welcomeMessage := fmt.Sprintf(":salt: Welcome <@%s>. Say hi!", m.User.ID)
    s.ChannelMessageSend(welcomeChannelID, welcomeMessage)
}
