package main

import (
	"log"

	"github.com/jemmycalak/mall-tangsel/server"
)

func main() {
	if err := server.Mains(); err != nil {
		log.Fatalln("app does'nt running !!")
		return
	}
	log.Println("App was running !")
}
