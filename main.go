package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

var lock bool = false

const MYID string = "385922547591675905"

func init() {
	Token = os.Getenv("DISCORD_TOKEN")
}

func main() {
	bot, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	bot.AddHandler(messageCreate)

	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.ID != MYID && lock == true {
		return
	}

	if m.Author.ID == MYID && strings.ToLower(m.Content) == "zenlock" {
		lock = true
	}

	if m.Author.ID == MYID && strings.ToLower(m.Content) == "zenunlock" {
		lock = false
	}

	// If the message is "ping" reply with "Pong!"
	if strings.ToLower(m.Content) == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if strings.ToLower(m.Content) == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if strings.ToLower(m.Content) == "ping pong" {
		s.ChannelMessageSend(m.ChannelID, "Pong! Ping!")
	}

	if m.Author.ID == MYID && strings.ToLower(m.Content) == "!luc" {
		s.ChannelMessageSend(m.ChannelID, "https://lh3.googleusercontent.com/pw/AIL4fc9kzv5TEvAwkO4SJcDbubGYlMfnK8KuxouW6NkC-lxHlDIJFbrLZ9wtvs94WwwdNj1ai6HmJW1BzjJvxK1l2_lMX5ZyNRVOO9-j2Imxi6DxxyB_p54BUiewzAyoJbVWneB1aQcWFJKQ6zZgWO5cO6b8ZeWEOk6duuXNiWRiIpt0_izlhxAWud0SVyB2UXN0d-tuUqYxV9M86fqnvUnyZm7OVLyXDXiZ7Gv9LXNyM9RJ05xRlHEN7jdSqQ_kBtuwB0MnuyJs_uYVzOWbxdzDwbApLoYH4l2mmb4iMAOgL1z1_MzD6HK2OPM5kbko9FDvkeymr2cWOAPYS2yayFSvvoCOIiRmfZVVq4oGRffSSGHeCV6GKhwlX6oUk2wCGN6GXIqSMObvgMd6vUi46OscDWRktrYitMiM7b35-QxlSM60hNbS22ObilMJZqOqKKTKCdAcdegjkmolopK9zeGV85jy9n5QyE_Jx3BheSdG0z9Q5Id81GkNRghGgD-GWWLTJrcuKe6Jg4NNAf0PylIJQeoqsngMg8Gfup4clwM4MhLoaWffRj6Kp_i7n2b510vTQCkIH14lUg6yaMdp9r2Xb81WVrWRagYWZAN1afqRHHpyqcvqooR-oNLfSHm6IX4d4fUKN-9LD7PBL7MsQpdvFZdCthHGQWGO9GIXuK29oLLsr7Vzow3r17WNW27fwbNq43Ne6UCxe0hKSV49CbDSeilvFcY8ezCSD3s97m_Yoa_GsWg2M4INlTbyIgihz4e4HRXMS0u8n_RhgMB4-IVHCeILnktz096MyDrI9hMLnvk6P8CBGArxBOXVVKYSyl8x_keAVLLQKYwEfFS6-14IhJw7n8ADlKZnyp3CTEOJChUzlkMLIsCoQ3KJ7LsZ6Gba1HY0hKIscxuYJAmCmTGC8g=w723-h1286-s-no?authuser=0")
	}

	if strings.ToLower(m.Content) == "!bored" {
		reply := generateRandomIdea()

		s.ChannelMessageSendReply(m.ChannelID, reply, m.Reference())
	}
}

func generateRandomIdea() string {
	games := []string{"Brick Rigs", "Minecraft Java", "Forza Horizon 4", "Forza Horizon 5", "Minecraft Bedrock", "No Man's Sky", "Gang Beast", "Roblox", "Rocket League", "Fall Guys", "Among Us"}
	ideas := []string{"Go Touch Grass", "Learn to cook", "Read a book", "Work out", "Watch Youtube", "Learn something new"}
	control1 := rand.Intn(1)

	var control string

	if control1 == 0 {
		control2 := rand.Intn(len(games))

		control = fmt.Sprintf("Why don't you play %s\n", games[control2])
	} else {
		control2 := rand.Intn(len(ideas))
		
		control = fmt.Sprintf("Why don't you %s\n", ideas[control2])
	}

	return control
}
