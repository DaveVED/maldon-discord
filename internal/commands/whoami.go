package commands

import (
    "fmt"

    "github.com/bwmarrin/discordgo"
)

func WhoamiMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    authorId := m.Author.ID
    authorUsername := m.Author.Username
    authorAvatar := m.Author.Avatar
    authorIconUrl := fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", authorId, authorAvatar)

    embed := &discordgo.MessageEmbed{
        Title: fmt.Sprintf("Check out %s's Maldon Profile!", authorUsername),
        Description: fmt.Sprintf("Curious about what Maldon knows about you %s? Use the `whoami` command to find out!", authorUsername),
        Color: 0x00ff00,
        Thumbnail: &discordgo.MessageEmbedThumbnail{
            URL: authorIconUrl,
        },
        Author: &discordgo.MessageEmbedAuthor{
            Name: authorUsername,
            URL: "https://maldon.herokuapp.com",
            IconURL: authorIconUrl,
        },
        Fields: []*discordgo.MessageEmbedField{
            {
                Name:   "Discord Username",
                Value:  authorUsername,
                Inline: true,
            },
            {
                Name: "Discord Discriminator",
                Value: m.Author.Discriminator,
                Inline: true,
            },
        },
    }
    s.ChannelMessageSendEmbed(m.ChannelID, embed)
}


