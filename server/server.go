package main

import (
	"encoding/json"
	"fmt"
	"github.com/aaspcodes/pingpong/shared"
	"net"
	"os"
	"sync"
)

func main() {
	fmt.Println("Starting the server...")
	Pong()
}

type client_map_struct struct {
	mu   sync.Mutex
	cmap map[string]int
}

func Pong() {
	// Set the port and start listening
	port := ":8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	client_map := client_map_struct{
		mu:   sync.Mutex{},
		cmap: make(map[string]int),
	}

	fmt.Println("Listening on " + port)

	for {
		// Wait for a connection.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			os.Exit(1)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections can be served concurrently.
		go handleRequest(conn, client_map)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, client_map client_map_struct) {
	// Make a buffer to hold incoming data.
	buffer := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
	}

	message := shared.Message_struct{}

	err = json.Unmarshal(buffer[:n], &message)

	if err != nil {
		fmt.Println(err, "server failed to unmarshal message")
	}

	// client_map.mu.Lock()
	// check if client is in map
	// if not, add client to map
	// check if new message is greater than current sequence number
	// if greater by 1, update sequence number
	// if greater by more than 1, send error message
	// if less than or equal to current sequence number, resend last message

	fmt.Println("Message received from client!: " + message.ToString() + "tiger")

	// Send a response back to the client contacting us.
	conn.Write([]byte("Message received from client!: " + string(buffer[:n]) + "dragon"))

	// Close the connection when you're done with it.
	conn.Close()
	fmt.Println("server finished receiving from client")
}
