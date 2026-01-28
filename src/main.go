package main

import (
	"log"
	"net/http"

	"github.com/FinnK04/bettercis-backend/modules"
	"github.com/FinnK04/bettercis-backend/modules/users"
)

func main() {
	mgr := modules.NewManager()

	// Register modules
	mgr.Register(users.New())
	// mgr.Register(products.New())  // Add more modules here

	log.Println("Server starting on :8080")
	log.Println("API available at http://localhost:8080/api")

	if err := http.ListenAndServe(":8080", mgr.Router()); err != nil {
		log.Fatal(err)
	}
}
