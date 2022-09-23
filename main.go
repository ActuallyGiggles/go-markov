package main

import (
	"log"
	"m/markov"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	instructions := markov.StartInstructions{
		Workers:       5,
		WriteInterval: 1,
		StartKey:      "start",
		EndKey:        "end",
	}

	markov.Start(instructions)

	<-sc
	log.Println("Stopping...")
}
