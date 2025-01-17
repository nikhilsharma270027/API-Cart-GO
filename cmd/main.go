package main

import (
	"log"

	"github.com/nikhilsharma270027/API-Cart-GO/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal("Server Running on port ")
	}
}
