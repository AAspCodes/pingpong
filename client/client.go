package main

import (
	"encoding/json"
	"fmt"
	"github.com/aaspcodes/pingpong/shared"
	"log"
	"net"
	"os"
	"time"
)

type client_struct struct {
	name            string
	sequence_number int
}

func main() {
	fmt.Println("Starting the client...")
	// Dial a TCP connection to localhost on port 8080

	client := client_struct{
		name:            os.Getenv("POD_NAME"),
		sequence_number: 0,
	}

	conn := create_conn()
	fmt.Println(client.name)

	message := shared.Message_struct{
		Sender:         client.name,
		SequenceNumber: client.sequence_number,
		Data:           "Hello from " + client.name,
	}

	send_message(conn, message)

	receive_message(conn)

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

func send_message(conn net.Conn, message shared.Message_struct) {
	// Message to send
	jsonData, err := json.Marshal(message.ToString())
	if err != nil {
		fmt.Println(err, "client failed to send message")
	}
	// Send the message
	_, err = conn.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}

func receive_message(conn net.Conn) {
	// Receive the response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ", string(buffer[:n]))
}
