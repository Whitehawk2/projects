package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get the client's IP address.  Handle X-Forwarded-For.
	ip := r.RemoteAddr
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ip = xForwardedFor // Use the first IP from X-Forwarded-For
		// In a production environment, you might want more sophisticated parsing
		// of X-Forwarded-For, as it can contain multiple comma-separated IPs.
	} else {
		// If X-Forwarded-For is not present, parse RemoteAddr to remove the port
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil { // Use the parsed host if no error
			ip = host
		}
	}
	log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, ip)

	// Write the response to the browser.
	fmt.Fprintf(w, "Hello %s\n", ip)
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portStr := os.Getenv("LISTEN_PORT")
	if portStr == "" {
		log.Fatal("LISTEN_PORT environment variable not set.")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %s", portStr)
	}
	if port < 1 || port > 65535 {
		log.Fatalf("Invalid port number: %d.  Port must be between 1 and 65535.", port)
	}

	// Register the handler function for the "/" path.
	http.HandleFunc("/", helloHandler)

	// Start the HTTP server.
	log.Printf("Listening on port %d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // Use ListenAndServe
	if err != nil {
		log.Fatal(err)
	}
}
