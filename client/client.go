package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting the client...")
	// Dial a TCP connection to localhost on port 8080
	conn := create_conn()

	pod_name := os.Getenv("POD_NAME")

	message := "hello from " + pod_name

	send_message(conn, message)

	conn.Close()
	time.Sleep(5 * time.Second)
	fmt.Println("client finished sending")
}

func create_conn() net.Conn {
	conn, err := net.Dial("tcp", "server-service:80")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func send_message(conn net.Conn, message string) {
	// Message to send
	jsonData, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err, "client failed to send message")
	}
	// Send the message
	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
