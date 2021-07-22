package main

import (
	"log"
	"net/http"
	"todo/route"
	"os"
	"todo/data"
)

func main() {
	log.Print("Server on")
	PORT := ":3000"

	data.InitDatabase()
	
	if envPort := os.Getenv("PORT"); envPort != "" {
		PORT = ":" + envPort
	}

	log.Print("THIS IS THE PORT ", PORT)
	http.ListenAndServe(PORT, route.RegisterRoute())
}
