package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	bot *discordgo.Session
)

func main() {

	err := godotenv.Load(".env")
	token := os.Getenv("DISCORD_BOT_TOKEN")

	if err != nil {
		fmt.Println("Error loading .env file!!!")
		return
	}

	bot, err = discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// botId = u.ID

	bot.AddHandler(ready)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	<-make(chan struct{})
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	log.Printf("%s %v", "momoko is ready, latency:", bot.HeartbeatLatency())
	s.UpdateGameStatus(0, "劍與魔法王國")
}
