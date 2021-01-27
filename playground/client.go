package main

import (
	"fmt"
	"net"
	// "time"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)


func main(){
	conn, _ := net.Dial(CONN_TYPE, CONN_HOST + ":"+CONN_PORT)

	go func(){
		for{
			var msg string
			fmt.Scan(&msg)
			fmt.Fprintf(conn,"%s\n", msg)
		}
	}()

	for{
		var msg string
		fmt.Fscan(conn, &msg)
		fmt.Println(msg)
	}
}