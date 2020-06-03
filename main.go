package main

import (
	"log"

	"github.com/AraqueGD/Twittor/bd"
	"github.com/AraqueGD/Twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Fail Connection to Data Base")
		return
	}
	handlers.Manage()
}
