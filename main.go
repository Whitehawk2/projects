package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Hello %s\n", conn.RemoteAddr().String())
	// You could add more logic here to interact with the client,
	// but for this example, we just print the greeting and close.
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port from the environment variable
	portStr := os.Getenv("LISTEN_PORT")
	if portStr == "" {
		log.Fatal("LISTEN_PORT environment variable not set.")
	}

	// Convert port string to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %s", portStr)
	}

	if port < 1 || port > 65535 {
		log.Fatalf("Invalid port number: %d.  Port must be between 1 and 65535.", port)
	}

	// Start listening on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Listening on port %d\n", port)

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue // Continue to the next connection attempt
		}
		go handleConnection(conn) // Handle the connection in a goroutine
	}
}
