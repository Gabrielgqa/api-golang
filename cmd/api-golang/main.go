package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gabrielgqa/api-golang/internal/data"
	"github.com/gabrielgqa/api-golang/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
	d.Close()
}
