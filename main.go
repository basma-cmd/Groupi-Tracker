package main

import (
	"log"
	"main/server"
	"os"
)

// Global variables to hold data fetched from the API

func main() {
	if len(os.Args) != 1 {
		log.Fatal("too many arguments!")
	}
	server.Server()
}
