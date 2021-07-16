package main

import (
	"log"
	"net/http"
	"todo/route"
	"os"
)

func main() {
	log.Print("Server on")
	PORT := ":3000"
	
	if envPort := os.Getenv("PORT"); envPort != "" {
		PORT = ":" + envPort
	}

	log.Print("THIS IS THE PORT ", PORT)
	http.ListenAndServe(PORT, route.RegisterRoute())
}
