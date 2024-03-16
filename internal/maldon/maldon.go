package maldon

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/DaveVED/maldon-discord/internal/handler"
    "github.com/bwmarrin/discordgo"
)

func Start() {
    // Log to indicate function start
    log.Println("Starting the bot...")

    token := os.Getenv("MALDON_DISCORD_BOT_TOKEN")

    // Log the token fetching
    log.Printf("Using token: %s\n", token)

    dg, err := discordgo.New("Bot " + token)
    if err != nil {
        log.Fatalf("Error creating Discord session: %v", err)
    }

    log.Println("Discord session created successfully.")

    dg.AddHandler(handler.MessageCreate)
    dg.Identify.Intents = discordgo.IntentsGuildMessages

    log.Println("Handler added.")

    if err = dg.Open(); err != nil {
        log.Fatalf("Error opening connection: %v", err)
    }

    log.Println("Bot is now running. Press CTRL+C to exit.")

    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    dg.Close()
}

