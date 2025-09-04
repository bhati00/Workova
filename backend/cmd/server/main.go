package server

import (
	"log"

	"github.com/bhati00/workova/backend/internal/bootstrap"
)

func main() {
	app := bootstrap.InitializeApp()
	app.Run(":8080")
	log.Println("Server running on http://localhost:8080")
}
