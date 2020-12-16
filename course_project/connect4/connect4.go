package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	CONN_HOST        = "localhost"
	CONN_PORT        = "12345"
	CONN_TYPE        = "tcp"
	BIG              = 100000
	SMALL            = -BIG
	PLAYER_ONE_COLOR = "○"
	PLAYER_TWO_COLOR = "◙"
	MIN_DIFFICULTY   = 1
	MAX_DIFFICULTY   = 10
)

var b *Board = NewBoard()

func init() {
	rand.Seed(time.Now().UnixNano())
}

func playAgainstAi() {

	fmt.Printf("Choose difficulty (number between %d and %d)", MIN_DIFFICULTY, MAX_DIFFICULTY)
	var option string
	fmt.Scan(&option)

	difficulty, err := strconv.Atoi(option)

	for err != nil || difficulty < MIN_DIFFICULTY || difficulty > MAX_DIFFICULTY {
		fmt.Println("Invalid input! Try again:")
		fmt.Scan(&option)
		difficulty, err = strconv.Atoi(option)
	}

	humanColor := PLAYER_ONE_COLOR
	aiColor := PLAYER_TWO_COLOR
	waiting := false

	for !b.areFourConnected(humanColor) && !b.areFourConnected(aiColor) {

		clearConsole()
		b.printBoard()

		if waiting {
			fmt.Println("waiting for oponent move...\n")
			_, bestMove := alphabeta(b, true, 0, SMALL, BIG, difficulty)
			b.drop(bestMove, aiColor)
			waiting = false
		} else {
			for {
				fmt.Printf("Enter column to drop: ")

				var column int
				_, err = fmt.Scan(&column)

				if err != nil || !b.drop(column, humanColor) {
					fmt.Println("You cant place here! Try another column")
				} else {
					waiting = true
					break
				}
			}
		}
	}

	clearConsole()
	b.printBoard()
	if b.areFourConnected(humanColor) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func playMultiplayer() {
	var conn net.Conn
	var color string
	var opponentColor string

	var waiting bool

	fmt.Println("Connecting to", CONN_TYPE, "server", CONN_HOST + ":" + CONN_PORT)
	conn, err := net.Dial(CONN_TYPE, CONN_HOST + ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	fmt.Println("searching for opponent...")

	var msg string
	fmt.Fscan(conn, &msg)

	if msg == "go" {
		color = PLAYER_ONE_COLOR
		opponentColor = PLAYER_TWO_COLOR
		waiting = false
	} else {
		color = PLAYER_TWO_COLOR
		opponentColor = PLAYER_ONE_COLOR
		waiting = true
	}

	for !b.areFourConnected(color) && !b.areFourConnected(opponentColor) {

		clearConsole()
		b.printBoard()

		if waiting {
			fmt.Println("waiting for oponent move...\n")

			c1 := make(chan int, 1)

			go func() {
				var column int
				fmt.Fscan(conn, &column)
				c1 <- column
			}()

			select {
			case otherPlayerColumn := <-c1:
				b.drop(otherPlayerColumn, opponentColor)
				waiting = false
			case <-time.After(60 * time.Second):
				fmt.Println("timeout Opponent failed to make a move in 60 seconds")
				return
			}

		} else {
			for {
				fmt.Printf("Enter column to drop: ")

				var column int
				_, err = fmt.Scan(&column)
				
				if err != nil || !b.drop(column, color) {
					fmt.Println("You cant place here! Try another column")
				} else {
					fmt.Fprintf(conn, "%d\n", column)
					waiting = true
					break
				}
			}
		}
	}

	fmt.Fprintf(conn, "end")

	clearConsole()
	b.printBoard()
	if b.areFourConnected(color) {
		fmt.Println("You won!")
	} else {
		fmt.Println("You lost.")
	}
}

func main() {

	fmt.Println("Hello! Welcome to connect four CMD!\n" +
		"To enter multiplayer lobby press [1]\n" + "To play against AI press [2]\n")

	var option string
	fmt.Scan(&option)

	for !(option == "1" || option == "2") {
		fmt.Println("Unknown command! Try again:")
		fmt.Scan(&option)
	}

	if option == "2" {
		playAgainstAi()
		return
	} else {
		playMultiplayer()
	}

}
