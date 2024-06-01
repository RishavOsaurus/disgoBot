package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	methods "disgoBot/methods"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	methods.Check(err)
	// Create a new ticker that triggers every "WAITTIME" seconds
	waittime, err := strconv.Atoi(os.Getenv("WAITTIME"))
	methods.Check(err)

	ticker := time.NewTicker(time.Duration(waittime) * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				publishedTitlesDiscord, publishedTitlesFacebook := methods.ReadJson()
				fetchedData := methods.Fetch();
				fmt.Println("Discord:", publishedTitlesDiscord)
				fmt.Println("FB: ", publishedTitlesFacebook)
				go methods.SendFacebook(fetchedData, publishedTitlesFacebook)
				go methods.SendEmbeds(fetchedData, publishedTitlesDiscord)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
