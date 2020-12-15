package main

import (
	"fmt"
	"net"
	"os"
	// "time"
)

// Application constants, defining host, port, and protocol.
const (
	CONN_HOST = "localhost"
	CONN_PORT = "12345"
	CONN_TYPE = "tcp"
)

var playerOne net.Conn

func main() {
	// Start the server and listen for incoming connections.
	fmt.Println("Starting " + CONN_TYPE + " server on " + CONN_HOST + ":" + CONN_PORT)
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		if playerOne == nil {
			playerOne = c
		} else {
			fmt.Fprint(c, "wait\n")
			fmt.Fprint(playerOne, "go\n")

			go handleConnection(playerOne, c)
			playerOne = nil
		}
	}
}

func handleConnection(conn1, conn2 net.Conn) {
	for {
		var msg string
		fmt.Fscan(conn1, &msg)
		fmt.Fprintf(conn2, "%s\n", msg)

		if msg == "end" {
			conn1.Close()
			conn2.Close()
			return
		}

		var msg1 string
		fmt.Fscan(conn2, &msg1)
		fmt.Fprintf(conn1, "%s\n", msg1)
	}
}
