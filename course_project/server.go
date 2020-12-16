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

var playerOneConn net.Conn

func main() {
	// Start the server and listen for incoming connections.
	fmt.Println("Starting " + CONN_TYPE + " server on " + CONN_HOST + ":" + CONN_PORT)
	l, err := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")
		fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")

		if playerOneConn == nil {
			playerOneConn = conn
		} else {
			fmt.Fprint(conn, "wait\n")
			fmt.Fprint(playerOneConn, "go\n")

			go handleConnection(playerOneConn, conn)
			playerOneConn = nil
		}
	}
}

func readMsgAndSend(from, to net.Conn) bool{
	var msg string
	_, err := fmt.Fscan(from, &msg)
	if err != nil{
		fmt.Println("Client " + from.RemoteAddr().String() + " disconnected.")
		return false
	}
	fmt.Fprintf(to, "%s\n", msg)
	return true
}

func handleConnection(conn1, conn2 net.Conn) {
	defer conn1.Close()
	defer conn2.Close()

	for{
		if !readMsgAndSend(conn1, conn2){
			return
		}
		if !readMsgAndSend(conn2, conn1){
			return
		}
	}
}
