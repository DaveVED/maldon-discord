package commands

import (
    "fmt"
    "context"

    "github.com/bwmarrin/discordgo"
    "github.com/DaveVED/maldon-database/pkg/users"
)

// WelcomeNewMember sends a welcome message to new members.
func WelcomeNewMember(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

    // TODO: Move the client Creation to application startup. This is for testing.
    // We should also move the channleID to a config file or something, along with the mongoURI.
    // We should still sned the weomoc message if we can't add the user to the database.
    mongoURI := "mongodb://localhost:27017"
    userInterface := users.NewClient(mongoURI)
    ctx := context.Background()
    
    // Create the users before send back to channel.
    user := users.DiscordUser{
        DiscordUsername: m.User.Username,
        DiscordDiscrim: m.User.Discriminator,
    }
    
    result, err := userInterface.CreateUser(ctx, user)
    if err != nil {
        fmt.Println(err)
    }
    
    welcomeChannelID := "1218320382717595748"
    welcomeMessage := fmt.Sprintf(":salt: Welcome <@%s>. Say hi!", m.User.ID)
    s.ChannelMessageSend(welcomeChannelID, welcomeMessage)
}
