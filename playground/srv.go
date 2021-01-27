package main

import (
	"fmt"
	"net"
	// "time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "12345"
	CONN_TYPE = "tcp"
)

func main(){
	
	l, _ := net.Listen(CONN_TYPE, ":" + CONN_PORT)
	
	var conn net.Conn
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	fmt.Println("Client connected.")
	fmt.Println("Client " + conn.RemoteAddr().String() + " connected.")

	go func(){
		for{
			var msg string
			fmt.Scan(&msg)
			fmt.Fprintf(conn, "%s\n", msg)
		}
	}()

	for{
		var msg string
		fmt.Fscan(conn, &msg)
		fmt.Println(msg)
	}
}