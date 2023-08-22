package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gabrielgqa/api-golang/internal/server"
)

func main() {
	serv, err := server.New("3000")
	if err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
}
