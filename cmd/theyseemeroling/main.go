package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/danmcclain/theyseemeroling/internal/lib"
)

func main() {
	discorder, err := lib.NewDiscorder(os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		log.Panicf("Bot go boom: %v", err)
	}

	defer discorder.Close()

	log.Println("up and running!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
