package main

import (
	// "encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting the server...")
	Pong()
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
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buffer := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
	}

	// var result shared.Message
	// jsonErr := json.Unmarshal(buffer[:n], &result)

	// if jsonErr != nil {
	// 	fmt.Println(jsonErr, "there")
	// }

	// message := fmt.Sprintf("pong got %s, %d", result.Msg, result.Num)
	// fmt.Println(message)
	fmt.Println(string(buffer[:n]))
	// // Send a response back to the person contacting us.
	// conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
	time.Sleep(5 * time.Second)
	fmt.Println("server finished receiving from client")
	// return message
}
