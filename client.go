package main

import (
	"fmt"
	"net"
	"os"
)

// Application constants, defining host, port, and protocol.
const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {

	// Start the client and connect to the server.
	fmt.Println("Connecting to", connType, "server", connHost+":"+connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	var msg string
	fmt.Fscan(conn, &msg)
	fmt.Println(msg)

	waiting := true

	if(msg == "go"){
		waiting = false
	}
	// run loop forever, until exit.
	for {

		if(!waiting){
			var text string
			fmt.Scan(&text)
			fmt.Fprintf(conn, "%s\n", text)
			waiting = true
		} else {
			var text string
			fmt.Fscan(conn, &text)
			fmt.Printf("%s\n", text)
			waiting = false
		}

	}
}