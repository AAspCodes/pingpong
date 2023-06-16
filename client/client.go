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
	conn, err := net.Dial("tcp", "server-service:80")
	if err != nil {
		log.Fatal(err)
	}

	pod_name := os.Getenv("POD_NAME")
	// Message to send
	jsonData, err := json.Marshal("hello from " + pod_name)
	if err != nil {
		fmt.Println(err, "here")
	}
	// Send the message
	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
	time.Sleep(5 * time.Second)
	fmt.Println("client finished sending")
}
