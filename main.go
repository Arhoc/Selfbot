package main

import (
	"github.com/bwmarrin/discordgo"

	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var (
	Token string = os.Getenv("TOKEN")
)

func main() {
	dg, err := discordgo.New("Bearer " + Token)
	if err != nil {
		panic(err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		panic(err)
	}

	uptimeCmd := exec.Command("uname -or")
	output, err := uptimeCmd.Output()
	if err != nil {
		panic(err)
	}

	err = dg.UpdateGameStatus(1, string(output))
	if err != nil {
		panic(err)
	}

	fmt.Println("Selfbot online.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		return
	}
}
