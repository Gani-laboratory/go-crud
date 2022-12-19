package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gani-laboratory/go-crud/config"
	"github.com/Gani-laboratory/go-crud/router"
)

func main() {
	config.CreateConnection()
	config.Migrate()
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server dijalankan pada port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}