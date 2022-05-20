package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if (host == "" || port == "") {
		log.Fatal("Missing environment variable DB_HOST or DB_PORT!")
	}
	
	fmt.Printf("Host: %s, Port: %s\n", host, port);
}
